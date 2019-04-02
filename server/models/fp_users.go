package models

import (
	"ForumPublica/server/db"
	"strings"
)

//User user model
type User struct {
	ID   int64  `gorm:"column:id;primary_key"`
	Role string `gorm:"column:role"`

	Characters []Character `gorm:"foreignkey:UserID"`
}

//TableName table name
func (u *User) TableName() string {
	return "fp_users"
}

//CharNames return user char names
func (u User) CharNames() string {
	var names []string
	db.DB.Model(&Character{}).Where("user_id = ?", u.ID).Order("name").Pluck("Name", &names)
	return strings.Join(names, ", ")
}
