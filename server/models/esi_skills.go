package models

//Skill skill model
type Skill struct {
	ID             int64  `gorm:"column:id;primary_key:yes"`
	EsiCharacterID int64  `gorm:"column:esi_character_id"`
	SkillID        int32  `gorm:"column:skill_id"`
	Level          int32  `gorm:"column:level"`
	Name           string `gorm:"column:name"`

	Character *Character `gorm:"foreignkey:EsiCharacterId"`
}

//TableName skill model table name
func (j *Skill) TableName() string {
	return "esi_skills"
}
