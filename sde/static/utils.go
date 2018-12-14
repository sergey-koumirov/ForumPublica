package static

import (
	"ForumPublica/sde/models"
	"math"
)

// TypeIdQuantity holds TypeId and Qty pairs
type TypeIdQuantity struct {
	TypeId   int32
	Quantity int64
}

// MaterialInfo holds type material description
type MaterialInfo struct {
	TypeId   int32
	Quantity int64
	HasBPO   bool
}

//IsT2BPO checks if given BPO is T2
func IsT2BPO(typeID int32) bool {
	_, exBpo := Blueprints[typeID]
	if !exBpo {
		return false
	}
	_, exBpo = T2toT1[typeID]
	return exBpo
}

//Level1BPOIds returns Level 1 components BPO IDs for given BPO
func Level1BPOIds(bpoID int32) []int32 {
	result := make([]int32, 0)
	for _, bpo := range Level1BPO(bpoID) {
		result = append(result, bpo.TypeId)
	}
	return result
}

//Level1BPO returns Level 1 components BPOs for given BPO
func Level1BPO(bpoID int32) []TypeIdQuantity {
	result := make([]TypeIdQuantity, 0)
	bpo := Blueprints[bpoID]

	if bpo.Manufacturing != nil {
		for _, mtr := range bpo.Manufacturing.Materials {
			bpoID, exists := BpoIdByTypeId[mtr.TypeId]
			if exists {
				result = append(result, TypeIdQuantity{TypeId: bpoID, Quantity: mtr.Quantity})
			}
		}
	}

	return result
}

//Level1Materials returns Level 1 (needed for immediate starting mnf job) materials/components for given BPO
func Level1Materials(bpoID int32) []MaterialInfo {
	result := make([]MaterialInfo, 0)
	bpo := Blueprints[bpoID]

	if bpo.Manufacturing != nil {
		for _, mtr := range bpo.Manufacturing.Materials {
			_, hasBPO := BpoIdByTypeId[mtr.TypeId]
			result = append(result, MaterialInfo{TypeId: mtr.TypeId, Quantity: mtr.Quantity, HasBPO: hasBPO})
		}
	}

	return result
}

func DefaultMeTe(bpoId int32) (int32, int32) {
	defaultME := int32(10)
	defaultTE := int32(20)
	if IsT2BPO(bpoId) {
		defaultME = int32(2)
		defaultTE = int32(4)
	}
	return defaultME, defaultTE
}

func ApplyME(repeats int64, cnt int64, me int32) int64 {
	return ApplyMEBonus(repeats, cnt, me, 0.0, 0.0)
}

func ApplyMEBonus(repeats int64, cnt int64, me int32, bonus1 float64, bonus2 float64) int64 {
	if cnt == 1 {
		return repeats
	}
	return int64(math.Ceil(float64(repeats*cnt) * (1.0 - float64(me)/100.0) * (1.0 - bonus1/100.0) * (1.0 - bonus2/100.0)))
}

func ProductIdByBpoId(bpoId int32) int32 {
	bpo := Blueprints[bpoId]
	if bpo.Manufacturing != nil {
		return bpo.Manufacturing.Products[0].TypeId
	} else {
		return 0
	}
}

func ProductByBpoId(bpoId int32) models.ZipType {
	return Types[ProductIdByBpoId(bpoId)]
}
