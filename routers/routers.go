package routers

import (
	"belajar-go-dasar/handler"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", handler.AppName)

	app.Post("/create", handler.Create)
	app.Get("/read", handler.Read)
	app.Put("/update", handler.Update)
	app.Delete("/delete", handler.Delete)
}
