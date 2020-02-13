package models

//TrRecord transaction info for index page
type TrRecord struct {
	ModelID       int64
	TypeID        int32
	TypeName      string
	Dt            string
	CharacterName string
	Quantity      int64
	Price         float64
	IsBuy         bool
	ClientName    string
	LocationName  string
	ImageURL      string
	InSummary     bool
}

//TrList list of transaction info for index page
type TrList struct {
	Records []TrRecord
	Page    int64
	Total   int64
	PerPage int64
}
