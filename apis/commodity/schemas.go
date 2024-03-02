package commodity

import "time"

type CreateCommodityRequest struct {
	DefaultName    string    `json:"default_name"`
	ProduceAt      time.Time `json:"produce_at"`
	ProduceAddress string    `json:"produce_address"`
	Category       string    `json:"category"`
}
