package models

import (
	"ForumPublica/sde/static"
)

//CnBlueprint model and calculated info about BPO
type CnBlueprint struct {
	Model         ConstructionBpo
	Runs          *[]ConstructionBpoRun
	IsGoal        bool
	IsT2          bool
	CopyTime      int64
	InventCnt     int64
	WholeCopyTime int64
	ReadyQty      int64
	InProdQty     int64
	DefaultME     int32
	PortionSize   int32
}

//CnBlueprints array of CnBlueprint
type CnBlueprints []CnBlueprint

func (s CnBlueprints) Len() int {
	return len(s)
}
func (s CnBlueprints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s CnBlueprints) Less(i, j int) bool {
	iID := s[i].Model.TypeId
	jID := s[j].Model.TypeId

	sameGoal := s[i].IsGoal == s[j].IsGoal && (static.Types[iID].Name < static.Types[jID].Name)
	diffGoal := s[i].IsGoal != s[j].IsGoal && s[i].IsGoal

	return sameGoal || diffGoal
}

//CnRecord bpo info for constructions page
type CnRecord struct {
	Model      Construction
	Title      string
	Blueprints CnBlueprints
	Components CnBlueprints
}

//CnList list of bpos info for constructions page
type CnList struct {
	Records []CnRecord
	Page    int64
	Total   int64
}
