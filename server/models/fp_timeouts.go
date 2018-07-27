package models

type Timeout struct {
	Key string `gorm:"column:skey;primary_key"`
	Dt  string `gorm:"column:dt"`
}

func (u *Timeout) TableName() string {
	return "fp_timeouts"
}
