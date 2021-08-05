package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/stevobengtson/budgetme_api/config"
	"github.com/stevobengtson/budgetme_api/middleware"
	"github.com/stevobengtson/budgetme_api/models"
	"github.com/stevobengtson/budgetme_api/routes"
)

func main() {
	var err error
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close() // Close the database connection when this function completes
	config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	middleware.SetupCors(r)
	middleware.SetupHealthcheck(r)
	r.Run(":8080")
}
