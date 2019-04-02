package models

//Location location names
type Location struct {
	ID   int64  `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}

//TableName table name
func (u *Location) TableName() string {
	return "esi_locations"
}
