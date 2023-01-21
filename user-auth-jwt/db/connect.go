package db

import (
	"fmt"
	"user-auth-jwt/config"
	"user-auth-jwt/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectToDatabase() {
	var err error

	port := config.Config("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Connected to Database")

	DB.AutoMigrate(&model.Product{}, &model.User{})
	fmt.Println("Database Migrated")

}