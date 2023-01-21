package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func UpdateTest(req *fiber.Ctx) error {

	inputPayload := struct {
		Username string `json:"Username"`
		City     string `json:"city"`
		Age      int    `json:"age"`
	}{}

	if err := req.BodyParser(&inputPayload); err != nil {
		return err
	}

	fmt.Println(inputPayload)

	req.Status(201)
	return req.JSON(fiber.Map{"success": true, "message": "app is healthy", "data": inputPayload})
}
