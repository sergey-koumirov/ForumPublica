package static

import (
	sdem "ForumPublica/sde/models"
	"math"
	"sort"
)

// TypeIDQuantityName holds TypeID and Qty and Name
type TypeIDQuantityName struct {
	TypeID   int32
	Quantity int64
	Name     string
}

//TypeIDQuantityNameList array of TypeIDQuantityName
type TypeIDQuantityNameList []TypeIDQuantityName

func (s TypeIDQuantityNameList) Len() int {
	return len(s)
}
func (s TypeIDQuantityNameList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s TypeIDQuantityNameList) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// TypeIDQuantity holds TypeDd and Qty pairs
type TypeIDQuantity struct {
	TypeID   int32
	Quantity int64
}

// MaterialInfo holds type material description
type MaterialInfo struct {
	TypeID   int32
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
		result = append(result, bpo.TypeID)
	}
	return result
}

//Level1BPO returns Level 1 components BPOs for given BPO
func Level1BPO(bpoID int32) []TypeIDQuantity {
	result := make([]TypeIDQuantity, 0)
	bpo := Blueprints[bpoID]

	if bpo.Manufacturing != nil {
		for _, mtr := range bpo.Manufacturing.Materials {
			bpoID, exists := BpoIDByTypeID[mtr.TypeID]
			if exists {
				result = append(result, TypeIDQuantity{TypeID: bpoID, Quantity: mtr.Quantity})
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
			_, hasBPO := BpoIDByTypeID[mtr.TypeID]
			result = append(result, MaterialInfo{TypeID: mtr.TypeID, Quantity: mtr.Quantity, HasBPO: hasBPO})
		}
	}

	return result
}

//DefaultMeTe get default ME & TE
func DefaultMeTe(bpoID int32) (int32, int32) {
	defaultME := int32(10)
	defaultTE := int32(20)
	if IsT2BPO(bpoID) {
		defaultME = int32(2)
		defaultTE = int32(4)
	}
	return defaultME, defaultTE
}

//ApplyME apply ME to manufacturing amount
func ApplyME(repeats int64, cnt int64, me int32) int64 {
	return ApplyMEBonus(repeats, cnt, me, 0.0, 0.0)
}

//ApplyMEBonus apply ME and bonuses to manufacturing amount
func ApplyMEBonus(repeats int64, cnt int64, me int32, bonus1 float64, bonus2 float64) int64 {
	if cnt == 1 {
		return repeats
	}
	return int64(math.Ceil(float64(repeats*cnt) * (1.0 - float64(me)/100.0) * (1.0 - bonus1/100.0) * (1.0 - bonus2/100.0)))
}

//ApplyTE apply TE to manufacturing time
func ApplyTE(seconds int64, te int32, skills []sdem.RawSkill) int64 {
	return ApplyTEBonus(seconds, te, 0.0, 0.0, 0.0, skills)
}

//ApplyTEBonus apply TE and bonuses to manufacturing time
func ApplyTEBonus(seconds int64, te int32, bonus1 float64, bonus2 float64, space float64, skills []sdem.RawSkill) int64 {
	citadelFactor := (1.0 - bonus1/100.0) * (1 - bonus2*space/100)
	skillFactor := (1.0 - 4*5/100.0) * (1.0 - 3*5/100.0)
	//3380 - Industry
	//3388 - Advanced Industry
	for _, sk := range skills {
		if sk.TypeID != 3380 && sk.TypeID != 3388 {
			skillFactor = skillFactor * (1.0 - 0.01*5)
		}
	}

	// toso use science skills
	teFactor := (1 - float64(te)/100.0)
	return int64(math.Ceil(float64(seconds) * teFactor * skillFactor * citadelFactor))
}

//ProductIDByBpoID get product id ny bpo id
func ProductIDByBpoID(bpoID int32) int32 {
	bpo := Blueprints[bpoID]
	if bpo.Manufacturing != nil && len(bpo.Manufacturing.Products) > 0 {
		return bpo.Manufacturing.Products[0].TypeID
	}
	return 0
}

//ProductByBpoID get product by bpo id
func ProductByBpoID(bpoID int32) sdem.ZipType {
	return Types[ProductIDByBpoID(bpoID)]
}

//MnfTime get manufacturing time from BPO
func MnfTime(bpoID int32) int32 {
	b, ex := Blueprints[bpoID]
	if ex {
		return b.Manufacturing.Time
	}
	return 0
}

//T1CopyTime T1 BPO 1 run copy time max skills
func T1CopyTime(bpoID int32) int32 {
	t1Id, existsT1 := T2toT1[bpoID]
	if existsT1 {
		t1Bpo := Blueprints[t1Id]
		if t1Bpo.Copying != nil {
			return t1Bpo.Copying.Time
		}
	}
	return 0
}

//InventTime get bpo invent time
func InventTime(bpoID int32) int32 {
	t1Id, existsT1 := T2toT1[bpoID]
	if existsT1 {
		t1Bpo, exists := Blueprints[t1Id]
		if exists {
			return t1Bpo.Invention.Time
		}
	}
	return 0
}

//ProductQuantity manufactoring result batch size
func ProductQuantity(bpoID int32) int64 {
	b, ex := Blueprints[bpoID]
	if ex && len(b.Manufacturing.Products) > 0 {
		return b.Manufacturing.Products[0].Quantity
	}
	return 0
}

//T2Runs get t2 bpc runs (invent with no decryptors)
func T2Runs(bpoID int32) int64 {
	t1Id, existsT1 := T2toT1[bpoID]
	if existsT1 {
		t1Bpo := Blueprints[t1Id]
		for _, p := range t1Bpo.Invention.Products {
			if p.TypeID == bpoID {
				return p.Quantity
			}
		}
	}
	return 0
}

//T2BaseChance invent chance - no skills
func T2BaseChance(bpoID int32) float64 {
	t1Id, existsT1 := T2toT1[bpoID]
	if existsT1 {
		t1Bpo := Blueprints[t1Id]
		for _, p := range t1Bpo.Invention.Products {
			if p.TypeID == bpoID {
				return p.Probability
			}
		}
	}
	return 0
}

//T2Chance invent chancewith max skills
func T2Chance(bpoID int32) float64 {
	return T2BaseChance(bpoID) * (1 + 5.0/40.0 + 5.0/30.0 + 5.0/30.0)
}

//InventCount get invent runs for desirible qty
func InventCount(bpoID int32, qty int64) int64 {
	return int64(
		math.Ceil(
			float64(qty) / (float64(ProductQuantity(bpoID)) * float64(T2Runs(bpoID)) * T2Chance(bpoID)),
		),
	)
}

//T1BPOTypeForT2 T1 BPO for T2
func T1BPOTypeForT2(bpoID int32) *sdem.ZipType {
	t1Id, existsT1 := T2toT1[bpoID]
	if existsT1 {
		t1Bpo := Types[t1Id]
		return &t1Bpo
	}
	return nil
}

//T1BPOForT2 T1 BPO for T2
func T1BPOForT2(bpoID int32) *sdem.ZipBlueprint {
	t1Id, existsT1 := T2toT1[bpoID]
	if existsT1 {
		t1Bpo := Blueprints[t1Id]
		return &t1Bpo
	}
	return nil
}

//T1DecryptorsForT2 for T2
func T1DecryptorsForT2(bpoID int32) *TypeIDQuantityNameList {
	t1Id, existsT1 := T2toT1[bpoID]
	if existsT1 {
		result := make(TypeIDQuantityNameList, 0)
		activity := *Blueprints[t1Id].Invention
		for _, el := range activity.Materials {
			result = append(result, TypeIDQuantityName{TypeID: el.TypeID, Quantity: el.Quantity, Name: Types[el.TypeID].Name})
		}
		sort.Sort(result)
		return &result
	}
	return nil
}
