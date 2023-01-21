package main

import (
	"log"
	"user-auth-jwt/config"
	"user-auth-jwt/db"
	"user-auth-jwt/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	db.ConnectToDatabase()

	router.Routing(app)
	
	port := config.Config("PORT")

	

	log.Fatal(app.Listen(":" + port))

}
