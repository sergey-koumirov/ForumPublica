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

//ZipGroup in-memory cache model
type ZipGroup struct {
	ID         int32
	CategoryID int32
	Name       string
	Published  bool
}

//ZipGroups in-memory cache model
type ZipGroups map[int32]ZipGroup

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
	RegionID  int64
	Security  float64
	RegionKey string
}

//ZipSolarSystemsList in-memory cache model
type ZipSolarSystemsList []ZipSolarSystem

//ZipSolarSystems in-memory cache model
type ZipSolarSystems map[int64]ZipSolarSystem
