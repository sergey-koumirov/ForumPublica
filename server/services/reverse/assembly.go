package reverse

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/models"
	"ForumPublica/server/services/reverse/jobruns"
	"ForumPublica/server/services/reverse/netsort"
	"math"
	"sort"
)

// Assembly get full assembly info using goal BPO and runs
func Assembly(cn *models.Construction) models.CnBlueprints {
	sortedIds := netsort.ArrangeBPO(&cn.Bpos)

	result := models.CnBlueprints{}
	fillResult(&result, sortedIds, &cn.Bpos)

	jobruns.SetJobRuns(&result, sortedIds, cn)

	calcRunTime(&result)
	calcSgtRunQty(&result)

	return result
}

func calcRunTime(result *models.CnBlueprints) {
	for i := range *result {
		bpo := (*result)[i]
		t := int64(0)
		for _, r := range *bpo.Runs {
			t = t + int64(r.Repeats)*static.ApplyTE(r.Qty*int64(static.MnfTime(bpo.Model.TypeID)), r.TE)
		}
		(*result)[i].MnfTime = t
	}
}

func calcSgtRunQty(result *models.CnBlueprints) {
	for i := range *result {
		bpo := (*result)[i]

		days := math.Ceil(float64(bpo.MnfTime) / float64(24*60*60))

		if int64(days) == 1 {
			(*result)[i].SgtRepeats = 1
		} else if bpo.Model.Qty%int64(days*10) == 0 {
			(*result)[i].SgtRepeats = int64(days)
		} else {
			(*result)[i].SgtRepeats = int64(days) - 1
		}

		qtyPerDay := int64(math.Ceil(float64(bpo.Model.Qty)/(10*days)) * 10)

		(*result)[i].SgtRunQty = qtyPerDay
	}
}

func fillResult(result *models.CnBlueprints, sortedIds []int32, bpos *models.ConstructionBpos) {
	for _, id := range sortedIds {

		isGoal := false
		qty := int64(0)
		for _, bpo := range *bpos {
			if bpo.TypeID == id {
				isGoal = true
				qty = qty + bpo.Qty //multiple bpo allowed
			}
		}
		defaultME, _ := static.DefaultMeTe(id)
		*result = append(
			*result,
			models.CnBlueprint{
				Model: models.ConstructionBpo{
					TypeID: id,
					Qty:    qty,
				},
				IsT2:        static.IsT2BPO(id),
				IsGoal:      isGoal,
				DefaultME:   defaultME,
				PortionSize: static.ProductByBpoID(id).PortionSize,
			},
		)
	}

	sort.Sort(*result)
}
