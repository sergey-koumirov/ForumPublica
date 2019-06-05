package tasks

import (
	"ForumPublica/server/db"
	"ForumPublica/server/esi"
	"ForumPublica/server/models"
	"ForumPublica/server/utils"
	"fmt"
)

type partInfo struct {
	amount   int64
	sequence string
}

//TaskTestMarket updates prices using ESI API
func TaskTestMarket(user models.User) error {

	var chars []models.Character
	db.DB.Where("user_id = ?", user.ID).Find(&chars)

	var orderIDs = make([]int64, 0)

	for _, char := range chars {
		charAPI := char.GetESI()

		data, charErr := charAPI.CharactersOrders(char.ID)
		if charErr != nil {
			fmt.Println("err: ", charErr)
		} else {
			for _, r := range data.R {
				orderIDs = append(orderIDs, r.OrderID)
			}
		}

		fmt.Println(char.ID, char.Name)
	}

	fmt.Println("orderIDs", orderIDs)

	api := esi.ESI{}

	result, err := api.MarketsOrdersAll(10000002, 34, "sell")

	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("Markets Orders")
		// fmt.Println("Pages: ", result.Pages)
		// fmt.Println("Expires: ", result.Expires)

		total := int64(0)
		accumulator := int64(0)
		sequence := "undecided"
		parts := make([]partInfo, 0)

		for _, record := range result {
			possibleSequence := ""
			if utils.Find(orderIDs, record.OrderID) > -1 {
				possibleSequence = "mine"
			} else {
				possibleSequence = "other"
			}

			if possibleSequence != sequence && sequence != "undecided" {
				parts = append(parts, partInfo{amount: accumulator, sequence: sequence})
				accumulator = 0
			}
			sequence = possibleSequence
			total = total + record.VolumeRemain
			accumulator = accumulator + record.VolumeRemain

			s := ""
			if utils.Find(orderIDs, record.OrderID) > -1 {
				s = "M"
			}
			fmt.Printf("L: %d, Price: %.2f, Vol: %d  %s\n", record.LocationID, record.Price, record.VolumeRemain, s)

		}
		if sequence != "undecided" {
			parts = append(parts, partInfo{amount: accumulator, sequence: sequence})
		}
		fmt.Println("------------------")

		for _, r := range parts {
			fmt.Printf("%s %d %.2f\n", r.sequence, r.amount, float64(r.amount)/float64(total)*100)
		}

	}

	return nil
}
