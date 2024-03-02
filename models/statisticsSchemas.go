package models

type FavoriteStatisticsResponse struct {
	CommodityItem *CommodityItem
	Count         int `json:"count"`
}

type FavoriteStatistics struct {
	Count           int `json:"count"`
	CommodityItemID int `json:"commodity_item_id"`
}

type PriceStatisticsResponse struct {
	CommodityItem CommodityItem
	PriceVariance float32 `json:"price_variance"`
}

//type PriceStatisticsResponseCurrent struct {
//	CommodityItem *CommodityItem
//}

type SummaryResponse struct {
	Seller         Seller
	SellerNum      int `json:"seller_num"`
	Commodity      Commodity
	CommodityNum   int `json:"commodity_num"`
	Platform       Platform
	PlatformNum    int    `json:"platform_num"`
	Category       string `json:"category"`
	CategoryNum    int    `json:"category_num"`
	FavoriteNum    int    `json:"favorite_num"`
	MessageNum     int    `json:"message_num"`
	PriceChangeNum int    `json:"price_change_num"`
}
