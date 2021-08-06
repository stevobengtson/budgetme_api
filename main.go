package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/stevobengtson/user_service/config"
	"github.com/stevobengtson/user_service/middleware"
	"github.com/stevobengtson/user_service/models"
	"github.com/stevobengtson/user_service/routes"
)

func main() {
	var err error
	godotenv.Load()

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close() // Close the database connection when this function completes
	config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	middleware.SetupCors(r)
	middleware.SetupHealthcheck(r)
	r.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("BUDGETME_USER_API_LISTEN_PORT")))
}
