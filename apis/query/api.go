package query

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
)

// GetFavoriteStatistics @GetFavoriteStatistics
// @Router /api/favorite/statistics [post]
// @Summary Get favorite statistics
// @Description Get favorite statistics
// @Tags Query
// @Accept json
// @Produce json
// @Param json body FavoriteStatisticsRequest true "json"
// @Success 200 {array} models.FavoriteStatisticsResponse
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func GetFavoriteStatistics(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser == nil || tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can get favorite statistics")
	}
	var favoriteStatisticsRequest FavoriteStatisticsRequest
	if err := c.BodyParser(&favoriteStatisticsRequest); err != nil {
		return common.BadRequest("Invalid request body")
	}
	if favoriteStatisticsRequest.All {
		res, err := GetFavoriteStatisticsAll()
		if err != nil {
			return err
		}
		return c.JSON(res)
	}

	if favoriteStatisticsRequest.AgeStart > favoriteStatisticsRequest.AgeEnd || favoriteStatisticsRequest.AgeStart < 0 {
		return common.BadRequest("Invalid time range")
	}
	res, err := GetFavoriteStatisticsSome(favoriteStatisticsRequest.Gender, favoriteStatisticsRequest.AgeStart, favoriteStatisticsRequest.AgeEnd)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

// GetPriceStatistics @GetPriceStatistics
// @Router /api/price/statistics [post]
// @Summary Get price statistics
// @Description Get price statistics
// @Tags Query
// @Accept json
// @Produce json
// @Param json body PriceStatisticsRequest true "json"
// @Success 200 {array} models.PriceStatisticsResponse
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
func GetPriceStatistics(c *fiber.Ctx) error {
	//tmpUser, err := GetGeneralUser(c)
	//if err != nil {
	//	return err
	//}
	//if tmpUser == nil || tmpUser.UserType != "admin" {
	//	return common.Forbidden("Only admin can get price statistics")
	//}
	var priceStatisticsRequest PriceStatisticsRequest
	if err := c.BodyParser(&priceStatisticsRequest); err != nil {
		return common.BadRequest("Invalid request body")
	}
	if priceStatisticsRequest.TimeStart.Time.After(priceStatisticsRequest.TimeEnd.Time) {
		return common.BadRequest("Invalid time range")
	}
	res, err := GetPriceStatisticsHistory(priceStatisticsRequest.CommodityID, priceStatisticsRequest.TimeStart.Time, priceStatisticsRequest.TimeEnd.Time)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

// GetPersonalSummary @GetPersonalSummary
// @Router /api/annual/summary [get]
// @Summary Get personal annual summary
// @Description Get personal annual summary
// @Tags Query
// @Accept json
// @Produce json
// @Success 200 {object} models.SummaryResponse
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func GetPersonalSummary(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser == nil || tmpUser.UserType != "user" {
		return common.Forbidden("Only user can get personal annul summary")
	}
	res, err := GetAnnualSummary(tmpUser.ID)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

// GetAllSummary @GetAllSummary
// @Router /api/annual/summary/all [get]
// @Summary Get all annual summary
// @Description Get all annual summary
// @Tags Query
// @Accept json
// @Produce json
// @Success 200 {object} models.SummaryResponse
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func GetAllSummary(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser == nil || tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can get all annul summary")
	}
	res, err := GetAnnualSummary(-1)
	if err != nil {
		return err
	}
	return c.JSON(res)
}
