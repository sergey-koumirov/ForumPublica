package models

type User struct {
	Id   int64  `gorm:"column:id;primary_key"`
	Role string `gorm:"column:role"`

	Characters []Character `gorm:"foreignkey:UserId"`
}

func (u *User) TableName() string {
	return "fp_users"
}
