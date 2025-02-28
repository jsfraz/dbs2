package main

import (
	"dbs2/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	postgres, err := gorm.Open(postgres.Open(fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PW"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DB"))), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Panicln(err)
	}
	err = postgres.AutoMigrate(
		&models.User{},
		&models.Author{},
		&models.Genre{},
		&models.Book{},
		&models.Review{},
		&models.Order{},
		&models.Discount{},
	)
	if err != nil {
		log.Panicln(err)
	}
}
