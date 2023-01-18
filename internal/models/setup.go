package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Cannt load Environment")
	}

	DBDriver := os.Getenv("DB_DRIVER")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv(("DB_PASSWORD"))
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")

	DBUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)

	DB, err = gorm.Open(DBDriver, DBUrl)

	if err != nil {
		fmt.Println("Cannot connect to database ", DBDriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", DBDriver)
	}
	DB.AutoMigrate(
		&User{},
		&Category{},
		&PaymentMethod{},
		&Product{},
		&Cart{},
		&Transaction{},
	)
}
