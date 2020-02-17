package models

//MiMarketVolume model
type MiMarketVolume struct {
	MarketItemID int64
	Dt           string
	Vol          int64
	IsMy         bool
}

//MiHist model
type MiHist struct {
	ID    int64 `json:"-"`
	Dt    string
	Price float64
}

//MiLocation model
type MiLocation struct {
	ID            int64
	Type          string
	Name          string
	CharacterName string
}

//MiStore model
type MiStore struct {
	ID            int64
	Type          string
	Name          string
	CharacterName string
	Qty           int64
}

//MiRecord market item info for index page
type MiRecord struct {
	ModelID     int64
	TypeID      int32
	TypeName    string
	MyPrice     float64
	MyVol       int64
	StoreVol    int64
	D90Vol      int64
	D90Data     []Tr90d
	LowestPrice float64
	LowestHist  []MiHist
	UnitPrice   float64
	BottomPrice float64
	Locations   []MiLocation
	Stores      []MiStore
}

//MiList list of market item info for index page
type MiList struct {
	Records    []MiRecord
	Characters []CharIDName
	Page       int64
	Total      int64
	PerPage    int64
}
