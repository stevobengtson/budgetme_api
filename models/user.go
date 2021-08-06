package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stevobengtson/user_service/config"
)

func GetAllUsersPaged(users *[]User, pagination *Pagination) (err error) {
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	if err = queryBuider.Model(&User{}).Find(&users).Error; err != nil {
		return err
	}
	return nil
}

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id uint) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByEmail ... Fetch only one user by Email
func GetUserByEmail(user *User, email string) (err error) {
	if err = config.DB.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
