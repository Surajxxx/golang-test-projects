package controller

import (
	"fmt"
	"time"
	"user-auth-jwt/config"
	"user-auth-jwt/db"
	"user-auth-jwt/helpers"
	"user-auth-jwt/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CreateUser(req *fiber.Ctx) error {

	type Response struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	var userModel = db.DB
	user := new(model.User)

	fmt.Println(user)

	if err := req.BodyParser(user); err != nil {
		return req.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	hashPassword, err := helpers.HashPassword(user.Password)

	if err != nil {
		return req.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	user.Password = hashPassword

	result := userModel.Create(&user)

	fmt.Print(result)

	if err := result.Error; err != nil {
		return req.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	finalData := Response{
		Username: user.Username,
		Email:    user.Email,
	}

	return req.JSON(fiber.Map{"status": "success", "message": "Created user", "data": finalData})

}

func Login(req *fiber.Ctx) error {

	type inputPayload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	input := new(inputPayload)

	if err := req.BodyParser(&input); err != nil {
		return req.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	fmt.Println(input)

	var userModel = db.DB

	var user model.User

	err := userModel.Where(&model.User{Username: input.Username}).Find(&user)

	if err.Error != nil {
		return req.Status(400).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	hashedPassword := user.Password
	userId := user.Id

	if helpers.ComparePassword(input.Password, hashedPassword) != nil {
		return req.Status(400).JSON(fiber.Map{"status": "error", "message": "incorrect credentials ", "data": nil})
	}

	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(config.Config("JWT_SECRET"))
	t, err1 := token.SignedString(secret)

	if err1 != nil {
		return req.Status(500).JSON(fiber.Map{"status": "error", "message": "server error ", "data": err1})
	}

	return req.Status(200).JSON(fiber.Map{"status": "success", "message": "login successful", "data": t})

}

func GetUser(req *fiber.Ctx) error{
	userId, err := req.ParamsInt("id")

	fmt.Println(req.Locals("userId"))

	if err != nil {
		return req.Status(400).JSON(fiber.Map{"status": "error", "message": "id is required", "data": nil})
	}

	headers := req.GetReqHeaders()

	token := headers["Authorization"]

	fmt.Println(userId)
	fmt.Println(token)

	userModel := db.DB

	var user model.User

	result := userModel.Where(&model.User{Id : uint(userId)}).Find(&user)

	if result.Error != nil {
		return req.Status(400).JSON(fiber.Map{"status": "error", "message": "user not found", "data": nil})
	}

	type Response struct {
		Username string  `json:"username"`
		Email string  `json:"email"`
	}

	finalResult := Response{
		Username : user.Username,
		Email:  user.Email,
	}

	fmt.Println(finalResult)

	return req.Status(200).JSON(fiber.Map{"status": "success", "message": "user details", "data": finalResult})
}
