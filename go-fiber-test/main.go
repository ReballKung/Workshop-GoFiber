package main

import (
	"go-fiber-test/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New() // * ใช้ libary (fiber)
	routes.InetRoutes(app)
	app.Listen(":3000")
}
