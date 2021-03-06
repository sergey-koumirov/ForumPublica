package models

//TrByType model
type TrByType struct {
	TypeID     int32
	TypeName   string
	TotalQty   int64
	TotalValue float64
}

//TrByDay model
type TrByDate struct {
	Dt         string
	TotalValue float64
}

//TrSummary model
type TrSummary struct {
	Total    float64
	Total1d  float64
	ByDate   []TrByDate
	ByType   []TrByType
	ByType1d []TrByType
}

//Tr90dSummary model
type Tr90dSummary struct {
	Total int64
	R     []Tr90d
}

//Tr90d 90 day data
type Tr90d struct {
	Id int64  `gorm:"column:id" json:"-"`
	D  string `gorm:"column:d"`
	Q  int64  `gorm:"column:q"`
}
type AnyArr []interface{}

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
