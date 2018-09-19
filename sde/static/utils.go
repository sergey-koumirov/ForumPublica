package static

func IsT2BPO(typeId int32) bool {

	_, exBpo := Blueprints[typeId]

	if !exBpo {
		return false
	}

	_, exBpo = T2toT1[typeId]

	return exBpo

}
