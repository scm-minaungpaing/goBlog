package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/scm-minaungpaing/goBlog/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
