package query

import "src/models"

type FavoriteStatisticsRequest struct {
	All      bool  `json:"all"`
	Gender   *bool `json:"gender"`
	AgeStart int   `json:"age_start"`
	AgeEnd   int   `json:"age_end"`
}

type PriceStatisticsRequest struct {
	CommodityID int           `json:"commodity_id" validate:"required"`
	TimeStart   models.MyTime `json:"time_start" validate:"required"`
	TimeEnd     models.MyTime `json:"time_end" validate:"required"`
}
