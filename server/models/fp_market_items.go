package models

import "ForumPublica/server/db"

//MarketItem monitored market item
type MarketItem struct {
	ID     int64 `gorm:"column:id;primary_key"`
	TypeID int32 `gorm:"column:type_id"`
	UserID int64 `gorm:"column:user_id"`

	Locations []MarketLocation `gorm:"foreignkey:MarketItemID"`
}

//TableName table name
func (m *MarketItem) TableName() string {
	return "fp_market_items"
}

//Delete delete model and children
func (m *MarketItem) Delete() {
	for _, l := range m.Locations {
		l.Delete()
	}
	db.DB.Delete(&m)
}

//MarketLocation market location info
type MarketLocation struct {
	ID           int64 `gorm:"column:id;primary_key"`
	MarketItemID int64 `gorm:"column:market_item_id"`
	MarketItem   *MarketItem

	LocationType string `gorm:"column:location_type"`
	LocationID   int64  `gorm:"column:location_id"`

	StoreLocationType string `gorm:"column:store_location_type"`
	StoreLocationID   int64  `gorm:"column:store_location_id"`
	StoreQty          int64  `gorm:"column:store_qty"`

	EsiCharacterID int64 `gorm:"column:esi_character_id"`

	Datas []MarketData `gorm:"foreignkey:MarketLocationID"`
}

//TableName table name
func (m *MarketLocation) TableName() string {
	return "fp_market_locations"
}

//Delete delete model and children
func (m *MarketLocation) Delete() {
	for _, d := range m.Datas {
		d.Delete()
	}
	db.DB.Delete(&m)
}

//MarketData market data for market location
type MarketData struct {
	ID               int64 `gorm:"column:id;primary_key"`
	MarketLocationID int64 `gorm:"column:market_location_id"`
	MarketLocation   *MarketLocation
	Dt               string `gorm:"column:dt"`

	SellVol         int64   `gorm:"column:sell_vol"`
	SoldVol         int64   `gorm:"column:sold_vol"`
	SellLowestPrice float64 `gorm:"column:sell_lowest_price"`

	BuyVol          int64   `gorm:"column:buy_vol"`
	BoughtVol       int64   `gorm:"column:bought_vol"`
	BuyHighestPrice float64 `gorm:"column:buy_highest_price"`

	Deciles []MarketDecile `gorm:"foreignkey:MarketDataID"`
}

//TableName table name
func (m *MarketData) TableName() string {
	return "fp_market_data"
}

//Delete delete model and children
func (m *MarketData) Delete() {
	for _, d := range m.Deciles {
		d.Delete()
	}
	db.DB.Delete(&m)
}

//MarketDecile deciles for market data
type MarketDecile struct {
	ID           int64 `gorm:"column:id;primary_key"`
	MarketDataID int64 `gorm:"column:market_data_id"`
	MarketData   *MarketData
	Decile       int32   `gorm:"column:decile"`
	Kind         string  `gorm:"column:kind"`
	AvgPrice     float64 `gorm:"column:average_price"`
	Vol          int64   `gorm:"column:decile_vol"`
}

//TableName table name
func (m *MarketDecile) TableName() string {
	return "fp_market_deciles"
}

//Delete delete model and children
func (m *MarketDecile) Delete() {
	db.DB.Delete(&m)
}
