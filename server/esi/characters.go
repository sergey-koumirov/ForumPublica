package esi

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//Skill model
type Skill struct {
	ActiveSkillLevel   int32 `json:"active_skill_level"`
	SkillID            int32 `json:"skill_id"`
	SkillpointsInSkill int64 `json:"skillpoints_in_skill"`
	TrainedSkillLevel  int32 `json:"trained_skill_level"`
}

//CharactersSkills model
type CharactersSkills struct {
	Skills  []Skill `json:"skills"`
	TotalSP int64   `json:"total_sp"`
	Error   string  `json:"error"`
}

//CharactersSkillsResponse model
type CharactersSkillsResponse struct {
	R       CharactersSkills
	Expires time.Time
}

//CharactersSkills get character skills
func (esi *ESI) CharactersSkills(characterID int64) (*CharactersSkillsResponse, error) {
	url := fmt.Sprintf("%s/characters/%d/skills/", ESIRootURL, characterID)
	result := CharactersSkills{}

	expires, _, err := auth("GET", esi.AccessToken, url, &result)
	if err != nil {
		return nil, err
	}
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	return &CharactersSkillsResponse{R: result, Expires: expires}, nil
}

//IDAndName model
type IDAndName struct {
	ID   int64  `json:"character_id"`
	Name string `json:"character_name"`
}

//CharactersNames get names by ids
func (esi *ESI) CharactersNames(charIds []int64) ([]IDAndName, error) {

	qIds := make([]string, len(charIds))
	for i, id := range charIds {
		qIds[i] = strconv.FormatInt(id, 10)
	}

	url := fmt.Sprintf("%s/characters/names/?character_ids=%s", ESIRootURL, strings.Join(qIds, "%2C"))
	var result []IDAndName

	err := get(url, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
