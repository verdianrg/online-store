package config

import (
	"fmt"
	"log"
	"online-store/gen/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	// DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panicf("[ERROR] db connection %v", err)
	}

	err = runMigration(db)
	if err != nil {
		log.Panicf("[ERROR] db connection %v", err)
	}

	return db
}

func runMigration(db *gorm.DB) error {
	log.Println("running auto migrate")
	err := db.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.Cart{},
		&models.History{},
	)

	return err
}
