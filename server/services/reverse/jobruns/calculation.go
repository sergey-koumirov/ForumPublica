package jobruns

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/models"
	"fmt"
	"math"
)

//SetJobRuns calculates runs for every given BPO using sorted array
func SetJobRuns(bpoInfos *models.CnBlueprints, sortedBpoIds []int32, cn *models.Construction) {

	for _, el := range *bpoInfos {
		fmt.Printf("%+v\n", el)
	}

	bpoInfosHash := getInfoHash(bpoInfos)

	allRunsHash := getRunsHash(&cn.Runs)

	for _, id := range sortedBpoIds {
		runs, _ := allRunsHash[id]
		addPhantomRun(bpoInfosHash[id].Model, &runs, cn)
		applyRuns(id, bpoInfosHash, &runs)
	}

}

func applyRuns(bpoId int32, bpoInfosHash map[int32]*models.CnBlueprint, runs *[]models.ConstructionBpoRun) {
	bpoInfosHash[bpoId].Runs = runs

	fmt.Printf("=== %s x %d\n", static.Types[bpoId].Name, bpoInfosHash[bpoId].Model.Qty)

	materialBpos := static.Level1BPO(bpoId)

	for _, run := range *runs {

		productQty := jobsQty(run.TypeId, int64(run.Repeats)*run.Qty)

		fmt.Println("__________", run.Repeats, run.Qty, productQty)

		for _, materialBpo := range materialBpos {
			materialQty := static.ApplyME(productQty, materialBpo.Quantity, run.ME)
			batchedQty := jobsQty(materialBpo.TypeId, materialQty) * int64(static.ProductByBpoId(materialBpo.TypeId).PortionSize)

			fmt.Println("______________________", static.Types[materialBpo.TypeId].Name, productQty, materialBpo.Quantity, run.ME, batchedQty)

			bpoInfosHash[materialBpo.TypeId].Model.Qty = bpoInfosHash[materialBpo.TypeId].Model.Qty + batchedQty
		}

		fmt.Printf("      > [%d] %d x %d = %d\n", run.ME, run.Qty, run.Repeats, run.Qty*int64(run.Repeats))
	}

}

func jobsQty(bpoId int32, resultQty int64) int64 {
	batchSize := static.ProductByBpoId(bpoId).PortionSize
	return int64(math.Ceil(float64(resultQty) / float64(batchSize)))
}

func addPhantomRun(model models.ConstructionBpo, runs *[]models.ConstructionBpoRun, cn *models.Construction) {
	runsQty := int64(0)
	for _, run := range *runs {
		runsQty = runsQty + run.Qty
	}

	defaultME, defaultTE := static.DefaultMeTe(model.TypeId)

	*runs = append(
		*runs,
		models.ConstructionBpoRun{
			TypeId:      model.TypeId,
			Repeats:     1,
			ME:          defaultME,
			TE:          defaultTE,
			Qty:         model.Qty - runsQty,
			CitadelType: cn.CitadelType,
			RigFactor:   cn.RigFactor,
			SpaceType:   cn.SpaceType,
		},
	)
}

func getRunsHash(allRuns *models.ConstructionBpoRuns) map[int32][]models.ConstructionBpoRun {
	runsHash := map[int32][]models.ConstructionBpoRun{}

	for i, run := range *allRuns {
		_, ex := runsHash[run.TypeId]
		if !ex {
			runsHash[run.TypeId] = []models.ConstructionBpoRun{}
		}
		runsHash[run.TypeId] = append(runsHash[run.TypeId], (*allRuns)[i])
	}

	return runsHash
}

func getInfoHash(bpoInfos *models.CnBlueprints) map[int32]*models.CnBlueprint {

	bpoInfosHash := map[int32]*models.CnBlueprint{}

	for i := range *bpoInfos {
		key := (*bpoInfos)[i].Model.TypeId
		bpoInfosHash[key] = &(*bpoInfos)[i]

	}
	return bpoInfosHash
}
