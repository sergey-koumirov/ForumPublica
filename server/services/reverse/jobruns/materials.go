package jobruns

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/models"
)

//RunsToMaterials calculates materials for all runs in construction
func RunsToMaterials(bpoInfos models.CnBlueprints) []models.CnMaterial {
	result := map[int32]models.CnMaterial{}

	for _, info := range bpoInfos {
		for _, m := range static.Level1Materials(info.Model.TypeId) {
			if !m.HasBPO {
				cnMat, ex := result[m.TypeId]
				if !ex {
					cnMat.Model = static.Types[m.TypeId]
					cnMat.Qty = int64(0)
					cnMat.Excluded = false
				}
				for _, run := range *info.Runs {
					cnMat.Qty = cnMat.Qty + static.ApplyME(int64(run.Repeats), run.Qty, run.ME)
				}
				result[m.TypeId] = cnMat
			}
		}

	}

	final := []models.CnMaterial{}

	for _, v := range result {
		v.Volume = float64(v.Qty) * float64(v.Model.Volume)
		final = append(final, v)
	}

	return final
}
