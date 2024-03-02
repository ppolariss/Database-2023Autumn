package priceChange

import (
	"src/models"
)

type GetPriceChangeModel struct {
	CommodityItemID int           `json:"commodity_item_id" validate:"required"`
	TimeStart       models.MyTime `json:"time_start" validate:"required"`
	TimeEnd         models.MyTime `json:"time_end" validate:"required"`
}

type UpdatePriceChangeModel struct {
	ID int `json:"id"`
	//CommodityItemID int     `json:"commodity_item_id"`
	NewPrice float32 `json:"new_price"`
}
