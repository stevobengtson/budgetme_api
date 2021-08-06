package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("BUDGETME_USER_API_DB_HOST"),
		Port:     os.Getenv("BUDGETME_USER_API_DB_PORT"),
		User:     os.Getenv("BUDGETME_USER_API_DB_USER"),
		Password: os.Getenv("BUDGETME_USER_API_DB_PASSWORD"),
		DBName:   os.Getenv("BUDGETME_USER_API_DB_NAME"),
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
