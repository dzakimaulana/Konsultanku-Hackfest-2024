package database

import (
	"fmt"
	"konsultanku-app/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var ApiKey string

func DatabaseConnection() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("Environment is not detected")
	}

	ApiKey = os.Getenv("API_KEY")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err.Error())
	}

	databaseError := database.AutoMigrate(
		&models.StudentProfile{},
		&models.Team{},
		&models.TeamComment{},
		&models.MseProfile{},
		&models.Problem{},
		&models.Collaboration{},
	)
	if databaseError != nil {
		panic(err.Error())
	}

	DB = database
}
