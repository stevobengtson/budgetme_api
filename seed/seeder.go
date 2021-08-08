package seed

import (
	"github.com/stevobengtson/budgetme/config"
	"github.com/stevobengtson/budgetme/models"
)

var users = []models.User{
	{
		Name:     "Steven Bengtson",
		Email:    "steven.bengtson@me.com",
		Password: "password",
	},
	{
		Name:     "Test User",
		Email:    "user@test.com",
		Password: "password",
	},
}

func Load() error {

	err := config.DB.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}

	err = config.DB.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}

	for i := range users {
		err = config.DB.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}

	return nil
}
