package repository

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stevobengtson/budgetme/config"
	"github.com/stevobengtson/budgetme/models"
)

func GetAllUsersPaged(users *[]models.User, pagination *models.Pagination) (err error) {
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	if err = queryBuider.Model(&models.User{}).Find(&users).Error; err != nil {
		return err
	}
	return nil
}

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]models.User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *models.User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *models.User, id uint) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByEmail ... Fetch only one user by Email
func GetUserByEmail(user *models.User, email string) (err error) {
	if err = config.DB.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *models.User, id string) (err error) {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *models.User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
