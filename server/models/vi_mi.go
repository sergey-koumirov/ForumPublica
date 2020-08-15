package models

//type MiVolumes model
type MiVolumes struct {
	Dt string
	VV []MiVolume
}

//MiVolume model
type MiVolume struct {
	MarketItemID int64  `json:"-"`
	Dt           string `json:"-"`
	Vol          int64
	M            int64
}

//MiHist model
type MiHist struct {
	ID int64 `json:"-"`
	D  string
	P  float64
}

//MiLocation model
type MiLocation struct {
	ID            int64
	LocationID    int64
	Type          string
	Name          string
	CharacterName string
	Expiration    string
	OrderExpired  bool
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
	VolumeHist  []MiVolumes
	Warnings    map[string]bool
}

//MiList list of market item info for index page
type MiList struct {
	Records    []MiRecord
	Characters []CharIDName
	Page       int64
	Total      int64
	PerPage    int64
}
