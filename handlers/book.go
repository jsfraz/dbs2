package handlers

import (
	"dbs2/database"
	"dbs2/models"
	"dbs2/utils"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Vytvoření knihy.
//
//	@param c
//	@param request
//	@return error
func CreateBook(c *gin.Context, request *models.CreateBook) (*models.Book, error) {
	// Validace narození autora
	date, err := utils.ParseISO8601String(request.Published)
	if err != nil {
		c.AbortWithStatus(400)
		return nil, fmt.Errorf("chyba parsování datumu vydání: %s", err)
	}
	// Kontrola zda existuje autor
	authorExists, err := database.AuthorExistsById(request.AuthorId)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	if !authorExists {
		c.AbortWithStatus(404)
		return nil, fmt.Errorf("author s ID %d neexistuje", request.AuthorId)
	}
	// Kontrola zda kniha existuje podle ISBN
	bookExists, err := database.BookExistsByIsbn(request.Isbn)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	if bookExists {
		c.AbortWithStatus(409)
		return nil, fmt.Errorf("kniha s ISBN %s již existuje", request.Isbn)
	}
	// Kontrola zda žánry existují
	for _, gId := range request.GenreIds {
		genreExists, err := database.GenreExistsById(gId)
		if err != nil {
			c.AbortWithStatus(500)
			return nil, err
		}
		if !genreExists {
			c.AbortWithStatus(404)
			return nil, fmt.Errorf("žánr s ID %d neexistuje", gId)
		}
	}
	genres, err := database.GetGenresByIds(request.GenreIds)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	// Vytvoření knihy
	book := models.NewBook(request.Name, request.AuthorId, request.Summary, request.Isbn, request.Price, *date, false, *genres)
	err = database.CreateBook(book)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return book, nil
}

// Nahrání obrázku knize.
//
//	@param c
func UploadBookImage(c *gin.Context) {
	// Získání souborů
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, gin.H{"error": "chyba parsování formuláře"})
		return
	}
	// ID knihy
	bookIdStr := form.Value["bookId"][0]
	bookId, _ := strconv.Atoi(bookIdStr)
	// Kontrola zda kniha existuje
	exists, err := database.BookExistsById(uint(bookId))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if !exists {
		c.JSON(404, gin.H{"error": fmt.Sprintf("kniha s ID %d neexistuje", bookId)})
		return
	}
	// Načtení knihy
	book, err := database.GetBookById(uint(bookId))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// Přílohy
	uploadedFile := form.File["image"][0]
	// Kontrola formátu
	if filepath.Ext(uploadedFile.Filename) != ".jpg" {
		c.JSON(400, gin.H{"error": "jsou povoleny pouze .jpg oubrázky"})
		return
	}
	// Uložení souboru
	err = c.SaveUploadedFile(uploadedFile, fmt.Sprintf("./uploads/books/%d%s", bookId, filepath.Ext(uploadedFile.Filename)))
	if err != nil {
		c.JSON(500, gin.H{"error": "nepodařilo se uložit soubor"})
		return
	}
	// Aktualizace knihy pokud neměla obrázek
	if book.HasImage {
		book.HasImage = true
		// Zde se nevracím error, musel by se ošetřovat obrázek
		database.UpdateBook(book)
	}
	c.Status(200)
}

// Vrátí všechny knihy.
//
//	@param c
//	@return *[]models.Book
//	@return error
func GetAllBooks(c *gin.Context) (*[]models.Book, error) {
	books, err := database.GetAllBooks()
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return books, nil
}

// Vrátí obrázek knihy.
//
//	@param c
func GetBookImage(c *gin.Context) {
	id := c.Param("id")
	bookId, _ := strconv.Atoi(id)
	// Kontrola zda kniha existuje
	exists, err := database.BookExistsById(uint(bookId))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if !exists {
		// Nastavení content-type
		c.Header("Content-Type", "image/jpeg")
		// Odeslání souboru
		c.File("./static/images/book_404.jpg")
		return
	}
	// Kontrola zda má kniha obrázek
	book, err := database.GetBookById(uint(bookId))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if !book.HasImage {
		// Nastavení content-type
		c.Header("Content-Type", "image/jpeg")
		// Odeslání souboru
		c.File("./static/images/book_404.jpg")
		return
	}
	// Cesta k adresáři s obrázky
	imagePath := filepath.Join(fmt.Sprintf("./uploads/books/%s.jpg", id))
	// Kontrola existence souboru
	_, err = os.Stat(imagePath)
	if err != nil {
		// Nastavení content-type
		c.Header("Content-Type", "image/jpeg")
		// Odeslání souboru
		c.File("./static/images/book_404.jpg")
		return
	}
	// Nastavení content-type
	c.Header("Content-Type", "image/jpeg")
	// Odeslání souboru
	c.File(imagePath)
}

// Vyhledávání knih podle různých kritérií.
//
//	@param c
//	@return *[]models.Book
//	@return error
func SearchBooks(c *gin.Context, request *models.SearchBooks) (*[]models.Book, error) {
	if err := request.Validate(); err != nil {
		c.AbortWithStatus(400)
		return nil, err
	}
	books, err := database.SearchBooks(request)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return books, nil
}

// Vrátí knihu podle ID.
//
//	@param c
//	@param request
//	@return *models.Book
//	@return error
func GetBookById(c *gin.Context, request *models.Id) (*models.Book, error) {
	book, err := database.GetBookById(request.Id)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return book, nil
}

// Vrátí true nebo false podle toho, zda je kniha v košíku.
//
//	@param c
//	@param request
//	@return *models.TrueFalse
//	@return error
func IsBookInCart(c *gin.Context, request *models.Id) (*models.TrueFalse, error) {
	// Načtení uživatele
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return nil, errors.New("uživatel není v kontextu")
	}
	bookInCart, err := database.IsBookInCart(request.Id, u.(*models.User).ID)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return models.NewTrueFalse(bookInCart), nil
}

// Aktualizace knihy.
//
//	@param c
//	@param request
//	@return error
func UpdateBook(c *gin.Context, request *models.UpdateBook) error {
	// Validace narození autora
	date, err := utils.ParseISO8601String(request.Published)
	if err != nil {
		c.AbortWithStatus(400)
		return fmt.Errorf("chyba parsování datumu vydání: %s", err)
	}
	// Kontrola zda existuje autor
	authorExists, err := database.AuthorExistsById(request.AuthorId)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	if !authorExists {
		c.AbortWithStatus(404)
		return fmt.Errorf("author s ID %d neexistuje", request.AuthorId)
	}
	// Kontrola zda existují žánry
	for _, gId := range request.GenreIds {
		genreExists, err := database.GenreExistsById(gId)
		if err != nil {
			c.AbortWithStatus(500)
			return err
		}
		if !genreExists {
			c.AbortWithStatus(404)
			return fmt.Errorf("žánr s ID %d neexistuje", gId)
		}
	}
	// Žánry
	genres, err := database.GetGenresByIds(request.GenreIds)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	// Aktualizace
	book, err := database.GetBookById(request.Id)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	book.Name = request.Name
	book.AuthorID = request.AuthorId
	book.Summary = request.Summary
	book.Isbn = request.Isbn
	book.Price = request.Price
	book.Published = *date
	book.Genres = *genres
	err = database.UpdateBook(book)
	if err != nil {
		c.AbortWithStatus(500)
		return err
	}
	return nil
}

// Vrátí true nebo false podle toho, zda je kniha v seznamu přání.
//
//	@param c
//	@param request
//	@return *models.TrueFalse
//	@return error
func IsBookInWishlist(c *gin.Context, request *models.Id) (*models.TrueFalse, error) {
	// Načtení uživatele
	u, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(500)
		return nil, errors.New("uživatel není v kontextu")
	}
	bookInWishlist, err := database.IsBookInWishlist(request.Id, u.(*models.User).ID)
	if err != nil {
		c.AbortWithStatus(500)
		return nil, err
	}
	return models.NewTrueFalse(bookInWishlist), nil
}
