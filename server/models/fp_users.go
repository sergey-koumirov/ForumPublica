package models

type User struct {
	Id   int64  `xorm:"autoincr pk"`
	Role string `xorm:"role"`
}

func (u *User) TableName() string {
	return "fp_users"
}
