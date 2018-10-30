package reverse

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/models"
	"ForumPublica/server/services/reverse/netsort"
	"sort"
)

// Assembly get full assembly info using goal BPO and runs
func Assembly(bpos *models.ConstructionBpos, runs *models.ConstructionBpoRuns) models.CnBlueprints {
	sortedIds := netsort.ArrangeBPO(bpos)

	result := models.CnBlueprints{}
	fillResult(&result, sortedIds, bpos)
	sort.Sort(result)

	return result
}

func fillResult(result *models.CnBlueprints, sortedIds []int32, bpos *models.ConstructionBpos) {
	for _, id := range sortedIds {

		isGoal := false
		qty := int64(0)
		for _, bpo := range *bpos {
			if bpo.TypeId == id {
				isGoal = true
				qty = qty + bpo.Qty //multiple bpo allowed
			}
		}

		*result = append(
			*result,
			models.CnBlueprint{
				Model: models.ConstructionBpo{
					TypeId: id,
					Qty:    qty,
				},
				IsT2:   static.IsT2BPO(id),
				IsGoal: isGoal,
			},
		)
	}
}
