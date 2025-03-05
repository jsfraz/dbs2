package main

import (
	"dbs2/models"
	"dbs2/routes"
	"dbs2/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Nastavení logu
	log.SetPrefix("dbs2: ")
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lmicroseconds)

	log.Println("Startuju...")

	// Načtení proměnných prostředí
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	singleton := utils.GetSingleton()
	singleton.Config = config

	// Připojení k databázi
	postgres, err := gorm.Open(postgres.Open(fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PW"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DB"))), &gorm.Config{Logger: logger.Default.LogMode(singleton.Config.GetGormLogLevel())})
	if err != nil {
		log.Panicln(err)
	}
	singleton.PostgresDb = *postgres

	// Migrace databázového schéma
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

	// HTTP server
	router, err := routes.NewRouter()
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
