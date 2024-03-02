package item

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
	"strings"
)

// GetAllCommodity @GetAllCommodity
// @Router /api/commodity/all [get]
// @Summary 获取所有商品
// @Description 获取所有商品
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {array} models.CommodityItem
// @Failure 403 {object} common.HttpError
func GetAllCommodity(c *fiber.Ctx) error {
	items, err := GetItems()
	if err != nil {
		return err
	}
	return c.JSON(&items)
}

// GetMyCommodity @GetMyCommodity
// @Router /api/commodity/data [get]
// @Summary 获取自己的商品
// @Description 获取自己的商品
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {array} models.CommodityItem
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func GetMyCommodity(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "seller" {
		return common.Forbidden("Only seller can get their items")
	}
	items, err := GetItemsBySellerID(tmpUser.ID)
	if err != nil {
		return err
	}
	return c.JSON(&items)
}

// SearchCommodity @SearchCommodity
// @Router /api/search [post]
// @Summary 搜索商品
// @Description 搜索商品
// @Tags Item
// @Accept json
// @Produce json
// @Param json body SearchQuery true "Range: name, category; Search: content to search"
// @Success 200 {array} models.CommodityItem
func SearchCommodity(c *fiber.Ctx) error {
	var searchQuery SearchQuery
	err := c.BodyParser(&searchQuery)
	if err != nil {
		return common.BadRequest("Invalid request body")
	}
	var items []CommodityItem
	switch strings.ToLower(searchQuery.Range) {
	case "name":
		if searchQuery.Accurate {
			items, err = GetItemsByName(searchQuery.Search)
		} else {
			items, err = GetItemsByNameFuzzy(searchQuery.Search)
		}
	case "category":
		if searchQuery.Accurate {
			items, err = GetItemsByCategory(searchQuery.Search)
		} else {
			items, err = GetItemsByCategoryFuzzy(searchQuery.Search)
		}
	default:
		return common.BadRequest("Invalid search range")
	}
	if err != nil {
		return err
	}
	return c.JSON(&items)
}

// AddCommodity @AddCommodity
// @Router /api/commodity/item [post]
// @Summary 添加商品
// @Description 添加商品
// @Tags Item
// @Accept json
// @Produce json
// @Param json body CreateItemModel true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func AddCommodity(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" && tmpUser.UserType != "seller" {
		return common.Forbidden("Only admin and seller can add items")
	}
	var createItemModel CreateItemModel
	err = c.BodyParser(&createItemModel)
	if err != nil {
		return common.BadRequest("Invalid request body")
	}
	// check if valid
	var commodityItem = CommodityItem{
		ItemName:    createItemModel.ItemName,
		Price:       createItemModel.Price,
		PlatformID:  createItemModel.PlatformID,
		CommodityID: createItemModel.CommodityID,
	}

	if tmpUser.UserType == "seller" {
		commodityItem.SellerID = tmpUser.ID
	} else {
		commodityItem.SellerID = createItemModel.SellerID
		if commodityItem.SellerID == 0 {
			return common.BadRequest("Invalid seller id")
		}
	}

	return commodityItem.Create()
}

// AddBatchCommodity @AddBatchCommodity
// @Router /api/commodity/item/batch [post]
// @Summary 批量添加商品
// @Description 批量添加商品
// @Tags Item
// @Accept json
// @Produce json
// @Param json body []CreateItemModel true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func AddBatchCommodity(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" && tmpUser.UserType != "seller" {
		return common.Forbidden("Only admin and seller can add items")
	}
	var createItemModel []CreateItemModel
	err = c.BodyParser(&createItemModel)
	if err != nil {
		return common.BadRequest("Invalid request body")
	}
	// check if valid
	var commodityItems []CommodityItem

	if tmpUser.UserType == "seller" {
		for _, item := range createItemModel {
			commodityItems = append(commodityItems, CommodityItem{
				ItemName:    item.ItemName,
				Price:       item.Price,
				SellerID:    tmpUser.ID,
				PlatformID:  item.PlatformID,
				CommodityID: item.CommodityID,
			})
		}
	} else {
		for _, item := range createItemModel {
			if item.SellerID == 0 {
				return common.BadRequest("Invalid seller id")
			}
			commodityItems = append(commodityItems, CommodityItem{
				ItemName:    item.ItemName,
				Price:       item.Price,
				SellerID:    item.SellerID,
				PlatformID:  item.PlatformID,
				CommodityID: item.CommodityID,
			})
		}
	}

	return CreateItems(commodityItems)
}

// UpdateCommodity @UpdateCommodity
// @Router /api/commodity/item [put]
// @Summary 更新商品
// @Description 更新商品
// @Tags Item
// @Accept json
// @Produce json
// @Param json body UpdateItemModel true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func UpdateCommodity(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" && tmpUser.UserType != "seller" {
		return common.Forbidden("Only admin and seller can update items")
	}
	var updateItemModel UpdateItemModel
	err = c.BodyParser(&updateItemModel)
	if err != nil {
		return common.BadRequest("Invalid request body")
	}
	var commodityItem = CommodityItem{
		ID:          updateItemModel.CommodityItemID,
		ItemName:    updateItemModel.ItemName,
		Price:       updateItemModel.Price,
		CommodityID: updateItemModel.CommodityID,
		PlatformID:  updateItemModel.PlatformID,
	}
	if commodityItem.Price > 0 {
		is, err := IsPriceChangeToday(commodityItem.ID)
		if err != nil {
			return err
		}
		if is {
			return common.Forbidden("Price has been changed today")
		}
	}
	if tmpUser.UserType == "seller" {
		commodityItem.SellerID = tmpUser.ID
	} else if commodityItem.SellerID != 0 {
		commodityItem.SellerID = updateItemModel.SellerID
	}
	return commodityItem.Update()
}

// DeleteCommodity @DeleteCommodity
// @Router /api/commodity/item [delete]
// @Summary 删除商品
// @Description 删除商品
// @Tags Item
// @Accept json
// @Produce json
// @Param id path string true "item id"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func DeleteCommodity(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" && tmpUser.UserType != "seller" {
		return common.Forbidden("Only admin and seller can delete items")
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	return DeleteItemByID(id)
}

// GetItemsByCommodity @GetItemsByCommodity
// @Router /api/commodity/item/{id} [get]
// @Summary 根据商品ID获取多个商品项
// @Description 根据商品ID获取多个商品项
// @Tags Item
// @Accept json
// @Produce json
// @Param id path string true "commodity_id"
// @Success 200 {array} models.CommodityItem
// @Failure 403 {object} common.HttpError
func GetItemsByCommodity(c *fiber.Ctx) error {
	commodityID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	items, err := GetItemsByCommodityID(commodityID)
	if err != nil {
		return err
	}
	return c.JSON(&items)
}
