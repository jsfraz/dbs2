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

// Zjistí zda uživatel existuje podle mailu.
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
func CreateUser(user *models.User) error {
	err := utils.GetSingleton().PostgresDb.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// Zjistí zda uživatel existuje podle ID.
//
//	@param id
//	@return bool
//	@return error
func UserExistsById(id uint) (bool, error) {
	var count int64
	err := utils.GetSingleton().PostgresDb.Model(&models.User{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

// Vrátí uživatele podle ID.
//
//	@param id
//	@return *models.User
//	@return error
func GetUserById(id uint) (*models.User, error) {
	var user models.User
	err := utils.GetSingleton().PostgresDb.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Vrátí všechny uživatele podle rolí.
//
//	@param roles
//	@return *[]models.User
//	@return error
func GetUsersByRole(roles []models.Role) (*[]models.User, error) {
	var users []models.User = []models.User{}
	err := utils.GetSingleton().PostgresDb.Model(&models.User{}).Where("role in ?", roles).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

// Aktualizace uživatele.
//
//	@param user
//	@return error
func UpdateUser(user *models.User) error {
	return utils.GetSingleton().PostgresDb.Save(user).Error
}

// Smazání uživatelů podle ID.
//
//	@param ids
//	@return error
func DeleteUsers(ids []uint) error {
	return utils.GetSingleton().PostgresDb.Where("id IN ?", ids).Delete(&models.User{}).Error
}
