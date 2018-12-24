package models

//User user model
type User struct {
	ID   int64  `gorm:"column:id;primary_key"`
	Role string `gorm:"column:role"`

	Characters []Character `gorm:"foreignkey:UserId"`
}

//TableName table name
func (u *User) TableName() string {
	return "fp_users"
}
