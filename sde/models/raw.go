package models

//RawType json-to-go model
type RawType struct {
	GroupID      int32             `yaml:"groupID"`
	Descriptions map[string]string `yaml:"description"`
	Names        map[string]string `yaml:"name"`
	PortionSize  int32             `yaml:"portionSize"`
	Published    bool              `yaml:"published"`
	RaceID       int32             `yaml:"raceID"`
	Volume       float32           `yaml:"volume"`
}

//RawTypes json-to-go model
type RawTypes map[int32]RawType

//RawSkill json-to-go model
type RawSkill struct {
	Level  int32 `yaml:"level"`
	TypeID int32 `yaml:"typeID"`
}

//RawMaterial json-to-go model
type RawMaterial struct {
	Quantity    int64   `yaml:"quantity"`
	TypeID      int32   `yaml:"typeID"`
	Probability float64 `yaml:"probability"`
}

//RawActivity json-to-go model
type RawActivity struct {
	Time      int32         `yaml:"time"`
	Materials []RawMaterial `yaml:"materials"`
	Products  []RawMaterial `yaml:"products"`
	Skills    []RawSkill    `yaml:"skills"`
}

//RawBlueprint json-to-go model
type RawBlueprint struct {
	Activities         map[string]RawActivity `yaml:"activities"`
	BlueprintTypeID    int32                  `yaml:"blueprintTypeID"`
	MaxProductionLimit int32                  `yaml:"maxProductionLimit"`
}

//RawBlueprints json-to-go model
type RawBlueprints map[int32]RawBlueprint

//RawSolarSystem yaml-to-go model
type RawSolarSystem struct {
	Security          float64 `yaml:"security"`
	SolarSystemID     int64   `yaml:"solarSystemID"`
	SolarSystemNameID int64   `yaml:"solarSystemNameID"`
}

//RawRegion yaml-to-go model
type RawRegion struct {
	RegionID int64 `yaml:"regionID"`
}

//RawName yaml-to-go model
type RawName struct {
	ItemID   int64  `yaml:"itemID"`
	ItemName string `yaml:"itemName"`
}

//RawGroup yaml-to-go model
type RawGroup struct {
	CategoryID int32             `yaml:"categoryID"`
	Names      map[string]string `yaml:"name"`
	Published  bool              `yaml:"published"`
}

//RawGroups json-to-go model
type RawGroups map[int32]RawGroup
