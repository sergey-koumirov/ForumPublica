package jobruns

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/models"
	"sort"
)

//RunsToMaterials calculates materials for all runs in construction
func RunsToMaterials(bpoInfos models.CnBlueprints) []models.CnMaterial {
	result := map[int32]models.CnMaterial{}

	for _, info := range bpoInfos {
		// fmt.Println("---", info.Model.TypeName())
		for _, m := range static.Level1Materials(info.Model.TypeId) {
			// fmt.Println("--- ---", static.Types[m.TypeId].Name, m.Quantity, m.HasBPO)
			if !m.HasBPO {
				cnMat, ex := result[m.TypeId]
				if !ex {
					cnMat.Model = static.Types[m.TypeId]
					cnMat.Qty = int64(0)
					cnMat.Excluded = false
				}
				for _, run := range *info.Runs {
					cnMat.Qty = cnMat.Qty + int64(run.Repeats)*static.ApplyME(run.Qty, m.Quantity, run.ME)
					// fmt.Println("--- --- ---", int64(run.Repeats)*static.ApplyME(run.Qty, m.Quantity, run.ME))
				}
				result[m.TypeId] = cnMat
			}
		}

	}

	final := models.CnMaterials{}

	for _, v := range result {
		v.Volume = float64(v.Qty) * float64(v.Model.Volume)
		final = append(final, v)
	}

	sort.Sort(final)

	return final
}
