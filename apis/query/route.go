package query

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app fiber.Router) {
	app.Post("/favorite/statistics", GetFavoriteStatistics)
	app.Post("/price/statistics", GetPriceStatistics)
	app.Get("/annual/summary", GetPersonalSummary)
	app.Get("/annual/summary/all", GetAllSummary)
}
