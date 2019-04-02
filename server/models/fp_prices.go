package models

//Price price
type Price struct {
	ID        int64   `gorm:"column:id;primary_key"`
	TypeID    int32   `gorm:"column:type_id"`
	Source    string  `gorm:"column:source"`
	BuyPrice  float64 `gorm:"column:buy_price"`
	SellPrice float64 `gorm:"column:sell_price"`
	Dt        string  `gorm:"column:dt"`
	MarketVol int64   `gorm:"column:market_volume"`
}

//TableName table name
func (u *Price) TableName() string {
	return "fp_prices"
}
