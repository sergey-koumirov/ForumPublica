package tasks

import (
	"ForumPublica/server/db"
	"ForumPublica/server/esi"
	"ForumPublica/server/models"
	"fmt"
)

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

	result, err := api.MarketsOrders(10000002, 3090, "sell", 1)

	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("Markets Orders")
		fmt.Println("Pages: ", result.Pages)
		fmt.Println("Expires: ", result.Expires)

		total := int64(0)

		for _, record := range result.R {
			total = total + record.VolumeRemain
			fmt.Printf("%+v\n", record)
		}

		fmt.Println("% Vol: ", 490.0/float64(total)*100.0)
	}

	return nil
}
