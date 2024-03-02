package priceChange

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
)

// GetPriceChange @GetPriceChangeById
// @Router /api/price/history [post]
// @Summary Get priceChange by ID
// @Description Get priceChange by ID
// @Tags PriceChange
// @Accept json
// @Produce json
// @Param json body GetPriceChangeModel true "json"
// @Success 200 {array} models.PriceChange
func GetPriceChange(c *fiber.Ctx) error {
	var getPriceChangeModel GetPriceChangeModel
	if err := c.BodyParser(&getPriceChangeModel); err != nil {
		// If there's an error in parsing, return an error response
		return common.BadRequest("Invalid request body")
	}
	priceChanges, err := GetPriceChangeById(getPriceChangeModel.CommodityItemID, getPriceChangeModel.TimeStart.Time, getPriceChangeModel.TimeEnd.Time)
	if err != nil {
		return err
	}
	return c.JSON(&priceChanges)
}

// UpdatePriceChange @UpdatePriceChange
// @Router /api/price/history [put]
// @Summary Update priceChange
// @Description Update priceChange
// @Tags PriceChange
// @Accept json
// @Produce json
// @Param json body UpdatePriceChangeModel true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func UpdatePriceChange(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can update priceChange")
	}

	var updatePriceChangeModel UpdatePriceChangeModel
	if err := c.BodyParser(&updatePriceChangeModel); err != nil {
		return common.BadRequest("Invalid request body")
	}
	var priceChange = PriceChange{
		ID:       updatePriceChangeModel.ID,
		NewPrice: updatePriceChangeModel.NewPrice,
	}
	return priceChange.Update()
}

// AddBatchPriceChange @AddBatchPriceChange
// @Router /api/price/history/batch [post]
// @Summary Add batch priceChange
// @Description Add batch priceChange
// @Tags PriceChange
// @Accept json
// @Produce json
// @Param json body []models.PriceChange true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func AddBatchPriceChange(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can add priceChange")
	}

	var priceChanges []PriceChange
	if err := c.BodyParser(&priceChanges); err != nil {
		return common.BadRequest("Invalid request body")
	}

	return CreatePriceChanges(priceChanges)
}

// DeletePriceChange @DeletePriceChange
// @Router /api/price/history/{id} [delete]
// @Summary Delete priceChange by ID
// @Description Delete priceChange by ID
// @Tags PriceChange
// @Accept json
// @Produce json
// @Param id path int true "priceChange id"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func DeletePriceChange(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can delete priceChange")
	}

	priceChangeID, err := c.ParamsInt("id")
	if err != nil {
		return common.BadRequest("Invalid request body")
	}
	return DeletePriceChangeByID(priceChangeID)
}

// DeleteBatchPriceChange @DeleteBatchPriceChange
// @Router /api/item/history/{id} [delete]
// @Summary Delete batch priceChange by commodityItemID
// @Description Delete batch priceChange by commodityItemID
// @Tags PriceChange
// @Accept json
// @Produce json
// @Param id path int true "commodityItem id"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func DeleteBatchPriceChange(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can delete priceChange")
	}

	itemID, err := c.ParamsInt("id")
	if err != nil {
		return common.BadRequest("Invalid request body")
	}
	return DeletePriceChangeByCommodityItemID(itemID)
}
