package models

type Character struct {
	Id   int64  `xorm:"id"`
	Name string `xorm:"name"`
}

func (c *Character) TableName() string {
	return "esi_characters"
}
