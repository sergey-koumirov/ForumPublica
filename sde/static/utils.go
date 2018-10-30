package static

// TypeIdQuantity holds TypeId and Qty pairs
type TypeIdQuantity struct {
	TypeId   int32
	Quantity int64
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
func Level1Materials(bpoID int32) []TypeIdQuantity {
	result := make([]TypeIdQuantity, 0)
	bpo := Blueprints[bpoID]

	if bpo.Manufacturing != nil {
		for _, mtr := range bpo.Manufacturing.Materials {
			result = append(result, TypeIdQuantity{TypeId: mtr.TypeId, Quantity: mtr.Quantity})
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
