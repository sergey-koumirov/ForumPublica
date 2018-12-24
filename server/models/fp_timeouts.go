package models

//Timeout timeout model
type Timeout struct {
	Key string `gorm:"column:skey;primary_key"`
	Dt  string `gorm:"column:dt"`
}

//TableName table name
func (u *Timeout) TableName() string {
	return "fp_timeouts"
}
