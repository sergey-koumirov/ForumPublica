package models

//ZipType in-memory cache model
type ZipType struct {
	ID          int32
	GroupID     int32
	Name        string
	PortionSize int32
	Published   bool
	Volume      float32
}

//ZipTypes in-memory cache model
type ZipTypes map[int32]ZipType

//ZipBlueprint in-memory cache model
type ZipBlueprint struct {
	BlueprintTypeID    int32
	MaxProductionLimit int32
	Copying            *RawActivity
	Manufacturing      *RawActivity
	Invention          *RawActivity
	ResearchMaterial   *RawActivity
	ResearchTime       *RawActivity
}

//ZipBlueprints in-memory cache model
type ZipBlueprints map[int32]ZipBlueprint
