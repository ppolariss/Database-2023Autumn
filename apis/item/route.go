package item

import "github.com/gofiber/fiber/v2"

func RegisterRoutesWithoutAuthorization(app fiber.Router) {
	app.Get("/commodity/all", GetAllCommodity)
	app.Post("/search", SearchCommodity)
	app.Get("/item/commodity/:id", GetItemsByCommodity)
}

func RegisterRoutes(app fiber.Router) {
	app.Get("/commodity/data", GetMyCommodity)
	app.Post("/commodity/item", AddCommodity)
	app.Put("/commodity/item", UpdateCommodity)
	app.Delete("/commodity/item/:id", DeleteCommodity)
	app.Post("/commodity/item/batch", AddBatchCommodity)
}
