package orm

import "gorm.io/gorm"

type Matching struct {
	gorm.Model
	// UserIdBuy      uint   `json:"user_id_buy"`
	// UsernameBuy    string `json:"username_buy"`
	// UserIdSell     uint   `json:"user_id_sell"`
	// UsernameSell   string `json:"username_sell"`
	ProductIdBuy   uint   `json:"product_id_buy"`
	ProductIdSell  uint   `json:"product_id_sell"`
	MatchingStatus string `gorm:"default:available" json:"matching_status"`
}
