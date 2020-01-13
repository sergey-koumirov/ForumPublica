package tasks

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/services"
	"fmt"
)

//TaskUpdatePrices updates prices using ESI API
func TaskUpdatePrices() error {
	// fmt.Println("TaskUpdatePrices started", time.Now().Format("2006-01-02 15:04:05"))
	// fmt.Println("TaskUpdatePrices finished", time.Now().Format("2006-01-02 15:04:05"))

	for _, b := range static.Blueprints {
		if static.IsT2BPO(b.BlueprintTypeID) && static.Types[b.BlueprintTypeID].Published {

			qtyTotal := int64(1000)

			// fmt.Printf("[%d] %s\n", b.BlueprintTypeID, static.Types[b.BlueprintTypeID].Name)
			result := services.ConstructionByType(b.BlueprintTypeID, qtyTotal)

			mTotal := 0.0

			for _, m := range result.Materials {
				mTotal = mTotal + float64(m.Qty)*m.Price
			}

			iTotal := 0.0
			for _, b := range result.Blueprints {
				for _, d := range *b.T1Decryptors {
					iTotal = iTotal + float64(b.InventCnt)*float64(d.Quantity)*services.GetDefaultPrice(d.TypeID)
				}
			}

			uPrice := (iTotal + mTotal) * 1.05 / float64(qtyTotal)
			jPrice := services.GetDefaultPrice(static.ProductIDByBpoID(b.BlueprintTypeID))

			if jPrice/uPrice > 3 {
				// fmt.Println("-----------------------------")
				t := static.Types[b.BlueprintTypeID]
				g := static.Groups[t.GroupID]
				fmt.Printf("[%d - %s] %s ___ %f\n", b.BlueprintTypeID, g.Name, t.Name, jPrice/uPrice)
				// fmt.Println("Qty:", qtyTotal)
				// fmt.Printf("Materials: %f\n", mTotal)
				// fmt.Printf("Invent: %f\n", iTotal)
				// fmt.Printf("Jobs: %f\n", (iTotal+mTotal)*0.02)
				// fmt.Printf("Taxes: %f\n", (iTotal+mTotal)*0.03)
				// fmt.Printf("Unit price: %f\n", uPrice)
				// fmt.Printf("Jita price: %f\n", jPrice)
				// fmt.Printf("K: %f\n", jPrice/uPrice)
			}

			if jPrice/uPrice < 0.75 {
				t := static.Types[b.BlueprintTypeID]
				g := static.Groups[t.GroupID]
				fmt.Printf("                                     [%d - %s] %s ___ %f\n", b.BlueprintTypeID, g.Name, t.Name, jPrice/uPrice)
			}

		}

	}

	return nil
}
