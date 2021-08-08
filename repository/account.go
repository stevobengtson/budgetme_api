package repository

import (
	"github.com/stevobengtson/budgetme/config"
	"github.com/stevobengtson/budgetme/models"
)

func CreateAccount(account *models.Account) (err error) {
	if err = config.DB.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func GetAllAccounts(userId uint, account *[]models.Account) (err error) {
	if err = config.DB.Where("user_id = ?", userId).Find(account).Error; err != nil {
		return err
	}
	return nil
}
