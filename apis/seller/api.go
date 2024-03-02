package seller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
	"src/utils"
)

//// GetCurSeller @GetCurSeller
//// @Router /api/sellers/me [get]
//func GetCurSeller(ctx *fiber.Ctx) error {
//	return nil
//}

// GetSeller @GetSeller
// @Router /api/sellers/data [get]
// @Summary Get seller info
// @Description Get seller info
// @Tags Seller
// @Accept json
// @Produce json
// @Success 200 {object} models.Seller
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func GetSeller(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}

	if tmpUser.UserType != "seller" {
		return common.Forbidden("You are not a seller.")
	}

	seller, err := GetSellerByID(tmpUser.ID)
	if err != nil {
		return err
	}

	return c.JSON(&seller)
}

// GetAllSeller @GetAllSeller
// @Router /api/sellers/all [get]
// @Summary Get all sellers
// @Description Get all sellers
// @Tags Seller
// @Accept json
// @Produce json
// @Success 200 {array} models.Seller
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func GetAllSeller(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}

	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can get all sellers.")
	}

	sellers, err := GetSellers()
	if err != nil {
		return err
	}

	return c.JSON(&sellers)
}

// AddSeller @AddSeller
// @Router /api/sellers [post]
// @Summary Add seller
// @Description Add seller
// @Tags Seller
// @Accept json
// @Produce json
// @Param json body CreateSellerRequest true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func AddSeller(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can add seller")
	}

	var seller Seller
	err = common.ValidateBody(c, &seller)
	if err != nil {
		return err
	}
	seller.Password = utils.MakePassword(seller.Password)
	return seller.Create()
}

// UpdateSeller @UpdateSeller
// @Router /api/sellers [put]
// @Summary Update seller
// @Description Update seller
// @Tags Seller
// @Accept json
// @Produce json
// @Param json body models.Seller true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @Failure 404 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func UpdateSeller(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" && tmpUser.UserType != "seller" {
		return common.Forbidden("Only admin and seller can update seller")
	}
	var seller Seller
	err = common.ValidateBody(c, &seller)
	if err != nil {
		return err
	}
	if tmpUser.UserType == "seller" && tmpUser.ID != seller.ID {
		return common.Forbidden("You can only update your own info")
	}
	return seller.Update()
}

// DeleteSeller @DeleteSeller
// @Router /api/sellers [delete]
// @Summary Delete seller
// @Description Delete seller
// @Tags Seller
// @Accept json
// @Produce json
// @Param id path int true "seller id"
// @Success 200
// @Failure 403 {object} common.HttpError
// @Failure 404 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func DeleteSeller(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can delete seller")
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	return DeleteSellerByID(id)
}
