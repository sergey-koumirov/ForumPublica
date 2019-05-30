package models

//MiRecord market item info for index page
type MiRecord struct {
	Model    MarketItem
	TypeName string
}

//MiList list of market item info for index page
type MiList struct {
	Records []MiRecord
	Page    int64
	Total   int64
	PerPage int64
}
