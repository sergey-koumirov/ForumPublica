package models

type User struct {
	Id   int64  `xorm:"id"`
	Role string `xorm:"role"`
}

func (u *User) TableName() string {
	return "fp_users"
}
