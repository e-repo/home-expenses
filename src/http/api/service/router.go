package service

import "github.com/gofiber/fiber/v2"

func SetupRouters(app *fiber.App) {
	service := app.Group("/api/service")

	service.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Pong!")
	})
}
