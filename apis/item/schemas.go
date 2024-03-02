package item

type CreateItemModel struct {
	SellerID    int     `json:"seller_id"`
	CommodityID int     `json:"commodity_id" validate:"required"`
	PlatformID  int     `json:"platform_id" validate:"required"`
	Price       float32 `json:"price" validate:"required"`
	ItemName    string  `json:"item_name" validate:"required"`
}

//ProduceAt      time.Time `json:"produce_at" validate:"required"`
//ProduceAddress string    `json:"produce_address" validate:"required"`

type UpdateItemModel struct {
	SellerID        int     `json:"seller_id"`
	CommodityItemID int     `json:"commodity_item_id" validate:"required"`
	PlatformID      int     `json:"platform_id" validate:"required"`
	CommodityID     int     `json:"commodity_id" validate:"required"`
	ItemName        string  `json:"item_name" validate:"required"`
	Price           float32 `json:"price" validate:"required"`
}
