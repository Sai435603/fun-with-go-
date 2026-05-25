package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/sai435603/go-fiber-postgres/pkg/models"
)

var DbInstance *gorm.DB

func ConnectDB() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host,
		port,
		user,
		dbname,
		password,
		sslmode,
	)

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to the postgres db successfully done!")
	db.AutoMigrate(&models.Stock{})
	log.Println("Stock Schema Created...")
	DbInstance = db
}

func GetDb() *gorm.DB {
	if DbInstance == nil {
		ConnectDB()
	}
	return DbInstance
}
