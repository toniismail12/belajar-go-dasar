package main

import (
	"belajar-go-dasar/config"
	"belajar-go-dasar/database"
	"belajar-go-dasar/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// database connect
	database.Connect()

	app := fiber.New()

	// set cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.AllowOrigins(),
		AllowHeaders:     config.AllowHeaders(),
		AllowMethods:     config.AllowMethods(),
		AllowCredentials: true,
	}))

	// router
	routers.Setup(app)

	// get port
	port := config.Env("PORT")

	// run
	app.Listen(":" + port)

}
