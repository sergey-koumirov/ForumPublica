package models

//Client Client names model
type ClientName struct {
	ID   int64  `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}

//TableName transaction model table name
func (j *ClientName) TableName() string {
	return "esi_client_names"
}
