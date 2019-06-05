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

//ZipRegion in-memory cache model
type ZipRegion struct {
	ID        int64
	Name      string
	RegionKey string
}

//ZipRegions in-memory cache model
type ZipRegions []ZipRegion

//ZipSolarSystem in-memory cache model
type ZipSolarSystem struct {
	ID        int64
	Name      string
	Region    *ZipRegion
	Security  float64
	RegionKey string
}

//ZipSolarSystems in-memory cache model
type ZipSolarSystems []ZipSolarSystem
