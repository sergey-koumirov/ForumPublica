package jobruns

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/models"
	"fmt"
)

//RunsToMaterials calculates materials for all runs in construction
func RunsToMaterials(bpoInfos models.CnBlueprints) []models.CnMaterial {
	result := map[int32]models.CnMaterial{}

	for _, info := range bpoInfos {
		fmt.Println(info.Model.TypeName(), info.Model.Qty)

		for _, m := range static.Level1Materials(info.Model.TypeId) {
			fmt.Println("_ _ _ ", static.Types[m.TypeId].Name, m.Quantity, m.HasBPO)
			if !m.HasBPO {
				cnMat, ex := result[m.TypeId]
				if !ex {
					cnMat.Model = static.Types[m.TypeId]
					cnMat.Qty = int64(0)
					cnMat.Excluded = false
				}
				for _, run := range *info.Runs {
					fmt.Println(run.Repeats, run.Qty, info.PortionSize)

					cnMat.Qty = cnMat.Qty + static.ApplyME(int64(run.Repeats), run.Qty, run.ME)

				}
				result[m.TypeId] = cnMat
			}
		}

	}

	final := []models.CnMaterial{}

	for _, v := range result {
		final = append(final, v)
	}

	return final
}
