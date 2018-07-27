package models

type Skill struct {
	Id             int64  `gorm:"column:id;primary_key:yes"`
	EsiCharacterId int64  `gorm:"column:esi_character_id"`
	SkillId        int32  `gorm:"column:skill_id"`
	Level          int32  `gorm:"column:level"`
	Name           string `gorm:"column:name"`

	Character *Character `gorm:"foreignkey:EsiCharacterId"`
}

func (j *Skill) TableName() string {
	return "esi_skills"
}
