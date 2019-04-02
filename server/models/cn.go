package models

import (
	sdem "ForumPublica/sde/models"
	"ForumPublica/sde/static"
)

//CnBlueprint model and calculated info about BPO
type CnBlueprint struct {
	Model       ConstructionBpo
	Runs        *[]ConstructionBpoRun
	Expenses    []ConstructionExpense
	IsGoal      bool
	IsT2        bool
	CopyTime    int32
	InventTime  int32
	InventCnt   int64
	ReadyQty    int64
	InProdQty   int64
	DefaultME   int32
	PortionSize int32
	MnfTime     int64
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
	iID := s[i].Model.TypeID
	jID := s[j].Model.TypeID

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
	Materials  []CnMaterial
}

//CnList list of bpos info for constructions page
type CnList struct {
	Records []CnRecord
	Page    int64
	Total   int64
}

//CnMaterial materil info for constructions page
type CnMaterial struct {
	Model    sdem.ZipType
	Qty      int64
	Excluded bool
	Volume   float64
	Price    float64
}

//CnMaterials materils info for constructions page
type CnMaterials []CnMaterial

func (s CnMaterials) Len() int {
	return len(s)
}
func (s CnMaterials) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s CnMaterials) Less(i, j int) bool {
	im := s[i].Model
	jm := s[j].Model

	same := im.GroupID == jm.GroupID && (im.Name < jm.Name)
	diff := im.GroupID < jm.GroupID

	return same || diff
}
