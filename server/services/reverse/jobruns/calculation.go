package jobruns

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/models"
	"math"
)

//SetJobRuns calculates runs for every given BPO using sorted array
func SetJobRuns(bpoInfos *models.CnBlueprints, sortedBpoIds []int32, cn *models.Construction) {
	bpoInfosHash := getInfoHash(bpoInfos)

	allRunsHash := getRunsHash(&cn.Runs)

	for _, id := range sortedBpoIds {
		runs, _ := allRunsHash[id]
		addPhantomRun(bpoInfosHash[id].Model, &runs, cn)
		applyRunsMaterialsBpos(id, bpoInfosHash, &runs)
	}

}

func applyRunsMaterialsBpos(bpoID int32, bpoInfosHash map[int32]*models.CnBlueprint, runs *[]models.ConstructionBpoRun) {
	bpoInfosHash[bpoID].Runs = runs
	materialBpos := static.Level1BPO(bpoID)
	for _, run := range *runs {
		productQty := int64(run.Repeats) * run.Qty
		for _, materialBpo := range materialBpos {
			materialQty := static.ApplyME(productQty, materialBpo.Quantity, run.ME)
			bpoInfosHash[materialBpo.TypeID].Model.Qty = bpoInfosHash[materialBpo.TypeID].Model.Qty + materialQty
		}
	}
}

func jobsRunsCount(bpoID int32, resultQty int64) int64 {
	batchSize := static.ProductByBpoID(bpoID).PortionSize
	return int64(math.Ceil(float64(resultQty) / float64(batchSize)))
}

func addPhantomRun(cnBpo models.ConstructionBpo, runs *[]models.ConstructionBpoRun, cn *models.Construction) {
	runsQty := int64(0)
	for _, run := range *runs {
		runsQty = runsQty + run.Qty*int64(run.Repeats)
	}

	jobRunsCnt := jobsRunsCount(cnBpo.TypeID, cnBpo.Qty)

	defaultME, defaultTE := static.DefaultMeTe(cnBpo.TypeID)

	*runs = append(
		*runs,
		models.ConstructionBpoRun{
			TypeID:      cnBpo.TypeID,
			Repeats:     1,
			ME:          defaultME,
			TE:          defaultTE,
			ExactQty:    cnBpo.Qty - runsQty*int64(static.ProductByBpoID(cnBpo.TypeID).PortionSize),
			Qty:         jobRunsCnt - runsQty,
			CitadelType: cn.CitadelType,
			RigFactor:   cn.RigFactor,
			SpaceType:   cn.SpaceType,
		},
	)
}

func getRunsHash(allRuns *models.ConstructionBpoRuns) map[int32][]models.ConstructionBpoRun {
	runsHash := map[int32][]models.ConstructionBpoRun{}

	for i, run := range *allRuns {
		_, ex := runsHash[run.TypeID]
		if !ex {
			runsHash[run.TypeID] = []models.ConstructionBpoRun{}
		}
		runsHash[run.TypeID] = append(runsHash[run.TypeID], (*allRuns)[i])
	}

	return runsHash
}

func getInfoHash(bpoInfos *models.CnBlueprints) map[int32]*models.CnBlueprint {

	bpoInfosHash := map[int32]*models.CnBlueprint{}

	for i := range *bpoInfos {
		key := (*bpoInfos)[i].Model.TypeID
		bpoInfosHash[key] = &(*bpoInfos)[i]

	}
	return bpoInfosHash
}
