package database

import (
	"log"
	"os"

	"fiber-simple/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// connectDb
func ConnectDb() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.Cat{})

	// Seed data if table is empty
	var count int64
	db.Model(&models.Cat{}).Count(&count)
	if count == 0 {
		log.Println("Seeding data")

		cats := []models.Cat{
			{Name: "Felix", Breed: "Persian"},
			{Name: "Garfield", Breed: "Bengal"},
			{Name: "Tom", Breed: "Siamese"},
			{Name: "Çiko", Breed: "Sarman"},
			{Name: "Misket", Breed: "Van"},
			{Name: "Cırı", Breed: "Tekir"},
			{Name: "Köpük", Breed: "Siyam"},
			{Name: "Karamel", Breed: "Siyam"},
		}

		db.Create(&cats)
	}
	DB = Dbinstance{
		Db: db,
	}
}
