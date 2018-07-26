package models

// TYPES
type ZipType struct {
	Id          int32
	GroupId     int32
	Name        string
	PortionSize int32
	Published   bool
	Volume      float32
}

type ZipTypes map[int32]ZipType

// BLUEPRINTS
type ZipBlueprint struct {
	BlueprintTypeId    int32
	MaxProductionLimit int32

	Copying          *RawActivity
	Manufacturing    *RawActivity
	Invention        *RawActivity
	ResearchMaterial *RawActivity
	ResearchTime     *RawActivity
}

type ZipBlueprints map[int32]ZipBlueprint
