package commodity

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
	"strings"
)

// GetAllCommodity @GetAllCommodity
// @Router /api/commodities/all [get]
// @Summary 获取所有商品
// @Description 获取所有商品
// @Tags 商品
// @Accept json
// @Produce json
// @Success 200 {object} Commodity
func GetAllCommodity(c *fiber.Ctx) error {
	commodities, err := GetCommodities()
	if err != nil {
		return err
	}
	return c.JSON(&commodities)
}

// CreateCommodity @AddCommodity
// @Router /api/commodities [post]
// @Summary 添加商品
// @Description 添加商品
// @Tags 商品
// @Accept json
// @Produce json
// @Param json body CreateCommodityRequest true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func CreateCommodity(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden()
	}

	var commodity Commodity
	err = c.BodyParser(&commodity)
	if err != nil {
		return common.BadRequest("Invalid request body")
	}

	return commodity.Create()
}

// DeleteCommodity @DeleteCommodity
// @Router /api/commodities [delete]
// @Summary 删除商品
// @Description 删除商品
// @Tags 商品
// @Accept json
// @Produce json
// @Param id path string true "commodity id"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @Failure 404 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func DeleteCommodity(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden()
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return common.BadRequest("Invalid request body")
	}

	return DeleteCommodityByID(id)
}

// UpdateCommodity @UpdateCommodity
// @Router /api/commodities [put]
// @Summary 更新商品
// @Description 更新商品
// @Tags 商品
// @Accept json
// @Produce json
// @Param json body Commodity true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @Failure 404 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func UpdateCommodity(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden()
	}

	var commodity Commodity
	err = c.BodyParser(&commodity)
	if err != nil {
		return common.BadRequest("Invalid request body")
	}

	return commodity.Update()
}

// SearchCommodity @SearchCommodity
// @Router /api/commodities/search [post]
// @Summary 搜索商品
// @Description 搜索商品
// @Tags 商品
// @Accept json
// @Produce json
// @Param json body SearchQuery true "json"
// @Success 200 {array} models.Commodity
// @Failure 400 {object} common.HttpError
func SearchCommodity(c *fiber.Ctx) error {
	var searchQuery SearchQuery
	err := c.BodyParser(&searchQuery)
	if err != nil {
		return common.BadRequest("Invalid request body")
	}
	var commodities []Commodity
	switch strings.ToLower(searchQuery.Range) {
	case "name":
		if searchQuery.Accurate {
			commodities, err = GetCommoditiesByName(searchQuery.Search)
		} else {
			commodities, err = GetCommoditiesByNameFuzzy(searchQuery.Search)
		}
	case "category":
		if searchQuery.Accurate {
			commodities, err = GetCommoditiesByCategory(searchQuery.Search)
		} else {
			commodities, err = GetCommoditiesByCategoryFuzzy(searchQuery.Search)
		}
	default:
		return common.BadRequest("Invalid search range")
	}
	if err != nil {
		return err
	}
	return c.JSON(&commodities)
}
