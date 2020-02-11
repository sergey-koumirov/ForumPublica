package models

import "ForumPublica/server/db"

//MarketItem monitored market item
type MarketItem struct {
	ID     int64 `gorm:"column:id;primary_key"`
	TypeID int32 `gorm:"column:type_id"`
	UserID int64 `gorm:"column:user_id"`

	Locations []MarketLocation `gorm:"foreignkey:MarketItemID"`
	Stores    []MarketStore    `gorm:"foreignkey:MarketItemID"`
	Datas     []MarketData     `gorm:"foreignkey:MarketLocationID"`
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
	for _, s := range m.Stores {
		s.Delete()
	}
	for _, d := range m.Datas {
		d.Delete()
	}
	db.DB.Delete(&m)
}

//MarketLocation market location info
type MarketLocation struct {
	ID             int64  `gorm:"column:id;primary_key"`
	MarketItemID   int64  `gorm:"column:market_item_id"`
	LocationType   string `gorm:"column:location_type"`
	LocationID     int64  `gorm:"column:location_id"`
	EsiCharacterID int64  `gorm:"column:esi_character_id"`

	MarketItem *MarketItem
	Character  *Character `gorm:"foreignkey:EsiCharacterID"`
}

//TableName table name
func (m *MarketLocation) TableName() string {
	return "fp_market_locations"
}

//Delete delete model and children
func (m *MarketLocation) Delete() {
	db.DB.Delete(&m)
}

//MarketStore market store info
type MarketStore struct {
	ID             int64  `gorm:"column:id;primary_key"`
	MarketItemID   int64  `gorm:"column:market_item_id"`
	LocationType   string `gorm:"column:location_type"`
	LocationID     int64  `gorm:"column:location_id"`
	EsiCharacterID int64  `gorm:"column:esi_character_id"`
	StoreQty       int64  `gorm:"column:store_qty"`

	MarketItem *MarketItem
	Character  *Character `gorm:"foreignkey:EsiCharacterID"`
}

//TableName table name
func (m *MarketStore) TableName() string {
	return "fp_market_stores"
}

//Delete delete model and children
func (m *MarketStore) Delete() {
	db.DB.Delete(&m)
}

//MarketData market data for market location
type MarketData struct {
	ID              int64   `gorm:"column:id;primary_key"`
	MarketItemID    int64   `gorm:"column:market_item_id"`
	Dt              string  `gorm:"column:dt"`
	SellVol         int64   `gorm:"column:sell_vol"`
	SellLowestPrice float64 `gorm:"column:sell_lowest_price"`
	BuyVol          int64   `gorm:"column:buy_vol"`
	BuyHighestPrice float64 `gorm:"column:buy_highest_price"`
	MyVol           int64   `gorm:"column:my_vol"`
	MyLowestPrice   float64 `gorm:"column:my_lowest_price"`

	Screenshots []MarketScreenshot `gorm:"foreignkey:MarketLocationID"`
}

//TableName table name
func (m *MarketData) TableName() string {
	return "fp_market_data"
}

//Delete delete model and children
func (m *MarketData) Delete() {
	for _, d := range m.Screenshots {
		d.Delete()
	}
	db.DB.Delete(&m)
}

//MarketScreenshot deciles for market data
type MarketScreenshot struct {
	ID           int64   `gorm:"column:id;primary_key"`
	MarketDataID int64   `gorm:"column:market_data_id"`
	Vol          int64   `gorm:"column:vol"`
	Price        float64 `gorm:"column:price"`
	IsMy         bool    `gorm:"column:is_my"`
}

//TableName table name
func (m *MarketScreenshot) TableName() string {
	return "fp_market_screenshots"
}

//Delete delete model and children
func (m *MarketScreenshot) Delete() {
	db.DB.Delete(&m)
}
