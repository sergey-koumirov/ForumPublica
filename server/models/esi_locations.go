package models

//StationMinID minimal possible station ID
var StationMinID = int64(60_000_000)

//StationMaxID maximum possible station ID
var StationMaxID = int64(64_000_000)

//Location location names
type Location struct {
	ID            int64  `gorm:"column:id;primary_key"`
	Name          string `gorm:"column:name"`
	LastCheckAt   string `gorm:"column:last_check_at"`
	SolarSystemID int64  `gorm:"column:solar_system_id"`
	RegionID      int64  `gorm:"column:region_id"`
}

//TableName table name
func (u *Location) TableName() string {
	return "esi_locations"
}
