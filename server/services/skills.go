package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

func RefreshSkills(cid int64) {
	var char models.Character
	errSel := db.DB.Where("id = ?", cid).First(&char).Error
	if errSel != nil {
		fmt.Println("RefreshSkills errSel", errSel)
		return
	}

	api := char.GetESI()
	skills, errEsi := api.CharactersSkills(char.Id)
	if errEsi != nil {
		fmt.Println("RefreshSkills errEsi", errEsi)
		return
	}

	for _, skill := range skills.R.Skills {
		temp := models.Skill{
			EsiCharacterId: cid,
			SkillId:        skill.SkillId,
			Level:          skill.ActiveSkillLevel,
			Name:           static.Types[skill.SkillId].Name,
		}

		errSk := db.DB.Where("esi_character_id = ? and skill_id = ?", cid, temp.SkillId).First(&temp).Error
		if errSk == nil {
			db.DB.Model(&temp).Update("level", temp.Level)
		} else if errSk == gorm.ErrRecordNotFound {
			db.DB.Create(&temp)
		} else {
			fmt.Println("RefreshSkills errSk", errSk)
		}

	}

}
