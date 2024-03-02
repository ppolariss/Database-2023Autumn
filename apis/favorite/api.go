package favorite

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
)

// AddFavorite @AddFavorite
// @Router /api/favorites [post]
// @Summary AddFavorite
// @Description AddFavorite
// @Tags Favorite
// @Accept json
// @Produce json
// @Param json body AddFavoriteModel true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @Failure 404 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func AddFavorite(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "user" {
		return common.Forbidden("Only user can add favorites")
	}
	var AddFavoriteModel AddFavoriteModel
	err = c.BodyParser(&AddFavoriteModel)
	if err != nil {
		return common.BadRequest("Invalid request body")
	}
	var favorite = Favorite{
		UserID:          tmpUser.ID,
		CommodityItemID: AddFavoriteModel.ItemId,
		PriceLimit:      0,
	}
	return favorite.Create()
}

// AddCommodityFavorite @AddCommodityFavorite
// @Router /api/favorites/commodity [post]
// @Summary AddCommodityFavorite
// @Description AddCommodityFavorite
// @Tags Favorite
// @Accept json
// @Produce json
// @Param json body AddCommodityFavoriteModel true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @Failure 404 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func AddCommodityFavorite(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "user" {
		return common.Forbidden("Only user can add favorites")
	}
	var AddCommodityFavoriteModel AddCommodityFavoriteModel
	err = c.BodyParser(&AddCommodityFavoriteModel)
	if err != nil {
		return common.BadRequest("Invalid request body")
	}

	return CreateCommodityFavorite(AddCommodityFavoriteModel.CommodityId, tmpUser.ID)
}

// AddPriceLimit @AddPriceLimit
// @Router /api/price/limit [post]
// @Summary AddPriceLimit
// @Description AddPriceLimit
// @Tags Favorite
// @Accept json
// @Produce json
// @Param json body PriceLimitModel true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @Failure 404 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func AddPriceLimit(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "user" {
		return common.Forbidden("Only user can add price limit")
	}
	var PriceLimitModel PriceLimitModel
	err = c.BodyParser(&PriceLimitModel)
	if err != nil {
		return common.BadRequest("Invalid request body")
	}
	var favorite = Favorite{
		UserID:          tmpUser.ID,
		CommodityItemID: PriceLimitModel.ItemId,
		PriceLimit:      PriceLimitModel.PriceLimit,
	}
	return favorite.Update()
}

// GetAllFavorites @GetAllFavorites
// @Router /api/favorites/all [get]
// @Summary GetAllFavorites
// @Description GetAllFavorites
// @Tags Favorite
// @Accept json
// @Produce json
// @Success 200 {array} models.Favorite
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func GetAllFavorites(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}

	if tmpUser.UserType != "user" {
		return common.Forbidden("Only user can get favorites")
	}

	favorites, err := GetFavoritesByUserID(tmpUser.ID)
	if err != nil {
		return err
	}
	return c.JSON(&favorites)
}

// DeleteFavorite @DeleteFavorite
// @Router /api/favorites/{id} [delete]
// @Summary DeleteFavorite
// @Description DeleteFavorite
// @Tags Favorite
// @Accept json
// @Produce json
// @Param id path string true "item id"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 404 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func DeleteFavorite(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}

	if tmpUser.UserType != "user" {
		return common.Forbidden("Only user can delete favorites")
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var favorite = Favorite{
		UserID:          tmpUser.ID,
		CommodityItemID: id,
	}
	return favorite.Delete()
}
