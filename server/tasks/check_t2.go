package tasks

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/services"
	"ForumPublica/server/utils"
	"fmt"
	"time"
)

//TaskCheckT2 updates prices using ESI API
func TaskCheckT2() error {
	fmt.Println("TaskUpdatePrices started", time.Now().Format("2006-01-02 15:04:05"))
	above := make([]string, 0)
	below := make([]string, 0)

	for _, b := range static.Blueprints {
		t := static.Types[b.BlueprintTypeID]
		product := static.ProductByBpoID(b.BlueprintTypeID)

		if static.IsT2BPO(b.BlueprintTypeID) &&
			static.Types[b.BlueprintTypeID].Published &&
			utils.FindInt32([]int32{1707, 1708, 973}, t.GroupID) == -1 &&
			utils.FindInt32([]int32{}, product.GroupID) == -1 {

			uPrice := services.UnitPrice(b)
			jPrice := services.GetDefaultPrice(static.ProductIDByBpoID(b.BlueprintTypeID))

			if jPrice/uPrice > 3 {
				// fmt.Println("-----------------------------")
				g := static.Groups[t.GroupID]
				above = append(above, fmt.Sprintf("%10d | %-34s | %-60s | %6.2f", b.BlueprintTypeID, g.Name, t.Name, jPrice/uPrice))
			}

			if jPrice/uPrice < 0.75 && utils.FindInt32([]int32{166, 722, 725, 726, 727, 787, 1994}, t.GroupID) == -1 {
				t := static.Types[b.BlueprintTypeID]
				g := static.Groups[t.GroupID]
				below = append(below, fmt.Sprintf("%10d | %-34s | %-60s | %6.2f", b.BlueprintTypeID, g.Name, t.Name, jPrice/uPrice))
			}

		}

	}

	fmt.Println()
	fmt.Println("-----------ABOVE---------------")
	for _, s := range above {
		fmt.Println(s)
	}
	fmt.Println()

	fmt.Println("-----------BELOW---------------")
	for _, s := range below {
		fmt.Println(s)
	}
	fmt.Println()

	fmt.Println("TaskUpdatePrices finished", time.Now().Format("2006-01-02 15:04:05"))

	return nil
}
