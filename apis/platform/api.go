package platform

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
)

// GetAllPlatform @GetAllPlatform
// @Router /api/platform/all [get]
// @Summary Get all platform
// @Description Get all platform
// @Tags Platform
// @Accept json
// @Produce json
// @Success 200 {array} models.Platform
func GetAllPlatform(c *fiber.Ctx) error {
	platforms, err := GetPlatforms()
	if err != nil {
		return err
	}
	return c.JSON(&platforms)
}

// AddPlatform @AddPlatform
// @Router /api/platform [post]
// @Summary Add platform
// @Description Add platform
// @Tags Platform
// @Accept json
// @Produce json
// @Param json body CreatePlatformRequest true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func AddPlatform(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can add platform")
	}

	var platform Platform
	err = common.ValidateBody(c, &platform)
	if err != nil {
		return err
	}
	return platform.Create()
}

// UpdatePlatform @UpdatePlatform
// @Router /api/platform [put]
// @Summary Update platform
// @Description Update platform
// @Tags Platform
// @Accept json
// @Produce json
// @Param json body models.Platform true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func UpdatePlatform(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can update platform")
	}

	var platform Platform
	err = common.ValidateBody(c, &platform)
	if err != nil {
		return err
	}
	return platform.Update()
}

// DeletePlatform @DeletePlatform
// @Router /api/platform [delete]
// @Summary Delete platform
// @Description Delete platform
// @Tags Platform
// @Accept json
// @Produce json
// @Param id path int true "platform id"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func DeletePlatform(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can delete platform")
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return common.BadRequest("Invalid platform id")
	}
	return DeletePlatformByID(id)
}
