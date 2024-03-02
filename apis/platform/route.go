package platform

import "github.com/gofiber/fiber/v2"

func RegisterRoutesWithoutAuthorization(app fiber.Router) {
	app.Get("/platform/all", GetAllPlatform)
}

func RegisterRoutes(app fiber.Router) {
	app.Post("/platform", AddPlatform)
	app.Delete("/platform/:id", DeletePlatform)
	app.Put("/platform", UpdatePlatform)
}
