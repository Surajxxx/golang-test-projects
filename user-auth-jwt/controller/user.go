package controller

import (
	"fmt"
	"user-auth-jwt/db"
	"user-auth-jwt/helpers"
	"user-auth-jwt/model"

	"github.com/gofiber/fiber/v2"
)


func CreateUser (req *fiber.Ctx) error {
	
	type Response struct {
		Username string `json:"username"`
		Email string `json:"email"`
	}
	
	var userModel = db.DB
	user := new(model.User)

	fmt.Println(user)

	if err:= req.BodyParser(user); err != nil {
		return req.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	hashPassword, err := helpers.HashPassword(user.Password)

	if err != nil {
		return req.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	user.Password = hashPassword

	result := userModel.Create(&user)

	fmt.Print(result)

	if err:= result.Error; err != nil {
		return req.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	finalData := Response{
		Username: user.Username,
		Email: user.Email,
	}

	return req.JSON(fiber.Map{"status": "success", "message": "Created user", "data": finalData})

}