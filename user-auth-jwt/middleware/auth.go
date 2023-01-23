package middleware

import (
	"strings"
	"user-auth-jwt/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Authentication(req *fiber.Ctx) error {
	userId, err1 :=(req.ParamsInt("id"))

	if err1 != nil {
	
		return req.Status(400).JSON(fiber.Map{"status": "error", "message": "id is required", "data": nil})
	}

	id1 := float64(userId)

	headers := req.GetReqHeaders()

	token := strings.Split(headers["Authorization"], " ")[1]

	secretKey := []byte(config.Config("JWT_SECRET"))

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return req.Status(400).JSON(fiber.Map{"status": "error", "message": "malformed token", "data": err})
	}

	if !t.Valid {
		return req.Status(400).JSON(fiber.Map{"status": "error", "message": "invalid token", "data": nil})
	}


	claims, ok := t.Claims.(jwt.MapClaims); 

	if !ok {
		return req.Status(400).JSON(fiber.Map{"status": "error", "message": "invalid token", "data": nil})
	}
	id := claims["userId"]

	if id != id1 {
		return req.Status(403).JSON(fiber.Map{"status": "error", "message": "unauthorized", "data": nil})
	}

	req.Locals("userId", id)

 return req.Next()
}
