package favorite

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Get("/favorites/all", GetAllFavorites)
	app.Post("/favorites", AddFavorite)
	app.Post("/price/limit", AddPriceLimit)
	app.Delete("/favorites/:id", DeleteFavorite)
	app.Post("/favorites/commodity", AddCommodityFavorite)
}
