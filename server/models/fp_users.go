package models

type User struct {
	Id   int64  `gorm:"column:id;primary_key"`
	Role string `gorm:"column:role"`
}

func (u *User) TableName() string {
	return "fp_users"
}
