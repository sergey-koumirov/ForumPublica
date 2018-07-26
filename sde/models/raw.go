package models

// TYPES
type RawType struct {
	GroupID      int64             `yaml:"groupID"`
	Descriptions map[string]string `yaml:"description"`
	Names        map[string]string `yaml:"name"`
	PortionSize  int32             `yaml:"portionSize"`
	Published    bool              `yaml:"published"`
	RaceID       int32             `yaml:"raceID"`
	Volume       float32           `yaml:"volume"`
}

type RawTypes map[int32]RawType

// BLUEPRINTS
type RawSkill struct {
	Level  int32 `yaml:"level"`
	TypeID int32 `yaml:"typeID"`
}

type RawMaterial struct {
	Quantity int64 `yaml:"quantity"`
	TypeID   int32 `yaml:"typeID"`
}

type RawActivity struct {
	Time      int32         `yaml:"time"`
	Materials []RawMaterial `yaml:"materials"`
	Products  []RawMaterial `yaml:"products"`
	Skills    []RawSkill    `yaml:"skills"`
}

type RawBlueprint struct {
	Activities         map[string]RawActivity `yaml:"activities"`
	BlueprintTypeID    int32                  `yaml:"blueprintTypeID"`
	MaxProductionLimit int32                  `yaml:"maxProductionLimit"`
}

type RawBlueprints map[int32]RawBlueprint
