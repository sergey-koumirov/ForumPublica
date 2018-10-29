package models

type CnBlueprint struct {
	Model         ConstructionBpo
	IsT2          bool
	CopyTime      int64
	InventCnt     int64
	WholeCopyTime int64
}

type CnBlueprints []CnBlueprint

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
