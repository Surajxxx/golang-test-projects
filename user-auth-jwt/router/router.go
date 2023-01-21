package router

import (
	"user-auth-jwt/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Routing(app *fiber.App) {

	test := app.Group("/test", logger.New())

	test.Get("/", func (req *fiber.Ctx) error {
		req.Status(200)
		return req.JSON(fiber.Map{"success" : true, "message": "app is healthy", "data" : nil })
	})

	test.Post("/update", controller.UpdateTest)
	

	user := app.Group("/user", logger.New())

	user.Post("/create", controller.CreateUser)

}