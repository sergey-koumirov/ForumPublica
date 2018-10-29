package static

type TypeIdQuantity struct {
	TypeId   int32
	Quantity int64
}

func IsT2BPO(typeId int32) bool {

	_, exBpo := Blueprints[typeId]

	if !exBpo {
		return false
	}

	_, exBpo = T2toT1[typeId]

	return exBpo

}

func Level1BPOIds(bpoId int32) []int32 {
	result := make([]int32, 0)
	for _, bpo := range Level1BPO(bpoId) {
		result = append(result, bpo.TypeId)
	}
	return result
}

func Level1BPO(bpoId int32) []TypeIdQuantity {
	result := make([]TypeIdQuantity, 0)
	bpo := Blueprints[bpoId]

	if bpo.Manufacturing != nil {
		for _, mtr := range bpo.Manufacturing.Materials {
			bpoId, exists := BpoIdByTypeId[mtr.TypeId]
			if exists {
				result = append(result, TypeIdQuantity{TypeId: bpoId, Quantity: mtr.Quantity})
			}
		}
	}

	return result
}
