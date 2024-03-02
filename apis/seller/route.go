package seller

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/sellers/all", GetAllSeller)
	app.Get("/sellers/data", GetSeller)
	app.Post("/sellers", AddSeller)
	app.Delete("/sellers/:id", DeleteSeller)
	app.Put("/sellers", UpdateSeller)
	//app.Get("/seller/:id", GetSellerById)
	//app.Post("/seller", AddSeller)
	//app.Put("/seller/:id", UpdateSeller)
	//app.Delete("/seller/:id", DeleteSeller)
}
