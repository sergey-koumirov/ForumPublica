package esi

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ESISkill struct {
	ActiveSkillLevel   int32 `json:"active_skill_level"`
	SkillId            int32 `json:"skill_id"`
	SkillpointsInSkill int64 `json:"skillpoints_in_skill"`
	TrainedSkillLevel  int32 `json:"trained_skill_level"`
}

type ESICharactersSkills struct {
	Skills  []ESISkill `json:"skills"`
	TotalSP int64      `json:"total_sp"`
	Error   string     `json:"error"`
}

type CharactersSkillsResponse struct {
	R       ESICharactersSkills
	Expires time.Time
}

func (esi *ESI) CharactersSkills(characterId int64) (*CharactersSkillsResponse, error) {
	url := fmt.Sprintf("%s/characters/%d/skills/", ESI_ROOT_URL, characterId)
	result := ESICharactersSkills{}

	expires, _, err := auth("GET", esi.AccessToken, url, &result)
	if err != nil {
		return nil, err
	}
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	return &CharactersSkillsResponse{R: result, Expires: expires}, nil
}

type ESIIdAndName struct {
	Id   int64  `json:"character_id"`
	Name string `json:"character_name"`
}

func (esi *ESI) CharactersNames(charIds []int64) ([]ESIIdAndName, error) {

	qIds := make([]string, len(charIds))
	for i, id := range charIds {
		qIds[i] = strconv.FormatInt(id, 10)
	}

	url := fmt.Sprintf("%s/characters/names/?character_ids=%s", ESI_ROOT_URL, strings.Join(qIds, "%2C"))
	var result []ESIIdAndName

	err := get(url, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
