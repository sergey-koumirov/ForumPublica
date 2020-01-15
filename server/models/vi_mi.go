package models

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
	Model     MarketItem
	TypeName  string
	Locations []MiLocation
	Stores    []MiStore
}

//MiList list of market item info for index page
type MiList struct {
	Records    []MiRecord
	Characters []CharIDName
	Page       int64
	Total      int64
	PerPage    int64
}
