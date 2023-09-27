package main

import (
	"log"
	"mygoprogram/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setUpRoutes(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())
	app.Static("/", "./public")

	app.Get("/api/users", routes.GetAllUsers)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users/:id", routes.GetUserById)
	app.Delete("/api/users/:id", routes.DeleteUser)
	app.Put("/api/users/:id", routes.UpdateUser)
}

func main() {
	app := fiber.New(fiber.Config{})

	setUpRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	log.Fatal(app.Listen(":3000"))
}
