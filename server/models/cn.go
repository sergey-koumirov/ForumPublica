package models

import (
	"ForumPublica/sde/static"
)

type CnBlueprint struct {
	Model         ConstructionBpo
	IsGoal        bool
	IsT2          bool
	CopyTime      int64
	InventCnt     int64
	WholeCopyTime int64
}

type CnBlueprints []CnBlueprint

func (s CnBlueprints) Len() int {
	return len(s)
}
func (s CnBlueprints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s CnBlueprints) Less(i, j int) bool {
	iId := s[i].Model.TypeId
	jId := s[j].Model.TypeId

	sameGoal := s[i].IsGoal == s[j].IsGoal && (static.Types[iId].Name < static.Types[jId].Name)
	diffGoal := s[i].IsGoal != s[j].IsGoal && s[i].IsGoal

	return sameGoal || diffGoal
}

type CnRecord struct {
	Model      Construction
	Title      string
	Blueprints CnBlueprints
	Components CnBlueprints
}

type CnList struct {
	Records []CnRecord
	Page    int64
	Total   int64
}
