package reverse

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/models"
	"ForumPublica/server/services/reverse/netsort"
	"fmt"
)

func Assembly(
	bpos *models.ConstructionBpos,
	runs *models.ConstructionBpoRuns,
) models.CnBlueprints {
	// collect uniq bpos from construction tree
	// sort bpo network
	//

	result := make(models.CnBlueprints, 0)

	sortedIds := netsort.ArrangeBPO(bpos)

	for _, id := range sortedIds {
		fmt.Println(id, static.Types[id].Name)
	}

	return result
}
