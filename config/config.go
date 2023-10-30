package config

import (
	"fmt"
	"main/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	// config := map[string]string{
	// 	"DB_Username": "root",
	// 	"DB_Password": "Akunmysql1@",
	// 	"DB_Port":     "3306",
	// 	"DB_Host":     "127.0.0.1",
	// 	"DB_Name":     "quiz",
	// }

	connectionString :=
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Quiz{})
	DB.AutoMigrate(&models.Question{})
	DB.AutoMigrate(&models.UserQuizResult{})
}
