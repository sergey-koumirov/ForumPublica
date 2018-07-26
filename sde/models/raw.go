package models

// TYPES
type RawType struct {
	GroupId      int32             `yaml:"groupID"`
	Descriptions map[string]string `yaml:"description"`
	Names        map[string]string `yaml:"name"`
	PortionSize  int32             `yaml:"portionSize"`
	Published    bool              `yaml:"published"`
	RaceId       int32             `yaml:"raceID"`
	Volume       float32           `yaml:"volume"`
}

type RawTypes map[int32]RawType

// BLUEPRINTS
type RawSkill struct {
	Level  int32 `yaml:"level"`
	TypeId int32 `yaml:"typeID"`
}

type RawMaterial struct {
	Quantity int64 `yaml:"quantity"`
	TypeId   int32 `yaml:"typeID"`
}

type RawActivity struct {
	Time      int32         `yaml:"time"`
	Materials []RawMaterial `yaml:"materials"`
	Products  []RawMaterial `yaml:"products"`
	Skills    []RawSkill    `yaml:"skills"`
}

type RawBlueprint struct {
	Activities         map[string]RawActivity `yaml:"activities"`
	BlueprintTypeId    int32                  `yaml:"blueprintTypeID"`
	MaxProductionLimit int32                  `yaml:"maxProductionLimit"`
}

type RawBlueprints map[int32]RawBlueprint
