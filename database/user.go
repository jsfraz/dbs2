package database

import (
	"dbs2/models"
	"dbs2/utils"
)

// Vrátí uživatele podle mailu.
//
//	@param mail
//	@return *models.User
//	@return error
func GetUserByMail(mail string) (*models.User, error) {
	var user models.User
	err := utils.GetSingleton().PostgresDb.Model(&models.User{}).Where("mail = ?", mail).Attrs(models.User{}).FirstOrInit(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Zjistí zda uživatel existuje.
//
//	@param mail
//	@return bool
//	@return error
func UserExistsByMail(mail string) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.User{}).Where("mail = ?", mail).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

// Vytvoření uživatele.
//
//	@param user
//	@return error
func CreateUser(user models.User) error {
	err := utils.GetSingleton().PostgresDb.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
