package main

import (
	"fmt"
	"github.com/e-repo/expenses/cmd/server/env"
	"github.com/e-repo/expenses/src/http/api/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	env.Init()

	app := fiber.New()

	service.SetupRouters(app)

	err := app.Listen(":8000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
