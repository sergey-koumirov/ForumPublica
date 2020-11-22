package tasks

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
	"ForumPublica/server/utils"
	"fmt"
	"strings"
	"time"
)

//TaskCheckT2 updates prices using ESI API
func TaskCheckT2() error {
	fmt.Println("TaskUpdatePrices started", time.Now().Format("2006-01-02 15:04:05"))

	result := make([]models.Deviation, 0)

	for _, b := range static.Blueprints {
		t := static.Types[b.BlueprintTypeID]
		product := static.ProductByBpoID(b.BlueprintTypeID)

		if static.IsT2BPO(b.BlueprintTypeID) &&
			t.Published &&
			!strings.HasPrefix(t.Name, "Standup ") &&
			utils.FindInt32([]int32{1707, 1708, 973, 1992}, t.GroupID) == -1 &&
			utils.FindInt32([]int32{4067, 4065, 4060, 4061}, product.GroupID) == -1 {

			uPrice := services.UnitPrice(b)
			jPrice := services.GetDefaultPrice(static.ProductIDByBpoID(b.BlueprintTypeID))

			if jPrice/uPrice > 3 {
				result = append(result, models.Deviation{ID: b.BlueprintTypeID, K: jPrice / uPrice})
			}

			if jPrice/uPrice < 0.75 && utils.FindInt32([]int32{166, 722, 725, 726, 727, 787, 1994}, t.GroupID) == -1 {
				result = append(result, models.Deviation{ID: b.BlueprintTypeID, K: jPrice / uPrice})
			}

		}

	}

	db.DB.Delete(models.Deviation{})
	for _, s := range result {
		db.DB.Create(&s)
	}

	fmt.Println("TaskUpdatePrices finished", time.Now().Format("2006-01-02 15:04:05"))

	return nil
}
