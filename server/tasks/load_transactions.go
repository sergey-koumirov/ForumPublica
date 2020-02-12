package tasks

import (
	"ForumPublica/server/db"
	"ForumPublica/server/esi"
	"ForumPublica/server/models"
	"ForumPublica/server/utils"
	"fmt"
	"time"
)

type partInfo struct {
	amount   int64
	sequence string
}

//TaskLoadTransactions updates prices using ESI API
func TaskLoadTransactions() error {

	fmt.Println("TaskLoadTransactions started", time.Now().Format("2006-01-02 15:04:05"))

	var chars []models.Character
	db.DB.Find(&chars)

	for _, char := range chars {

		fmt.Println("   ", char.ID, char.Name)

		api, errApi := char.GetESI()
		if errApi != nil {
			fmt.Println("TaskLoadTransactions.api:", errApi)
		} else {

			var (
				errReq  error
				created = int(1)
				cnt     = int(1)
				r       *esi.CharactersWalletTransactionsResponse
				from_id = int64(0)
			)

			for cnt > 0 && float64(created)/float64(cnt) > 0.5 && errReq == nil {
				r, errReq = api.CharactersWalletTransactions(char.ID, from_id)
				fmt.Println("from_id", from_id)

				trIds := make([]int64, 0)
				for _, t := range r.R {
					trIds = append(trIds, t.TransactionId)
				}

				exTrIds := make([]int64, 0)
				db.DB.Model(&models.Transaction{}).Where("esi_character_id = ? and id in (?)", char.ID, trIds).Pluck("ID", &exTrIds)

				uniqClientIdsMap := make(map[int64]int32)

				for _, t := range r.R {
					if utils.FindInt64(exTrIds, t.TransactionId) == -1 {
						temp := models.Transaction{
							ID:             t.TransactionId,
							EsiCharacterID: char.ID,
							ClientID:       t.ClientId,
							Dt:             t.Date.Format("2006-01-02 15:04:05"),
							IsBuy:          t.IsBuy,
							IsPersonal:     t.IsPersonal,
							JournalRefID:   t.JournalRefId,
							LocationID:     t.LocationId,
							Quantity:       t.Quantity,
							TypeID:         t.TypeId,
							UnitPrice:      t.UnitPrice,
						}
						db.DB.Create(&temp)
						uniqClientIdsMap[t.ClientId] = 1
					}
				}

				uniqClietIds := make([]int64, 0)
				for k, _ := range uniqClientIdsMap {
					uniqClietIds = append(uniqClietIds, k)
				}

				exClIds := make([]int64, 0)
				db.DB.Model(&models.ClientName{}).Where("id in (?)", char.ID, uniqClietIds).Pluck("ID", &exClIds)

				norExClIds := utils.DiffInt64(uniqClietIds, exClIds)

				for _, batch := range utils.InBatchesInt64(norExClIds, 500) {
					names, errN := api.UniverseNames(batch)
					if errN != nil {
						fmt.Println("TaskLoadTransactions.errN:", errN)
					} else {
						for _, name := range names {
							temp := models.ClientName{
								ID:   name.ID,
								Name: name.Name,
							}
							db.DB.Create(&temp)
						}
					}
				}

				created = len(trIds) - len(exTrIds)
				cnt = len(r.R)
				if cnt > 0 {
					from_id = r.R[cnt-1].TransactionId
				}
			}

			if errReq != nil {
				fmt.Println("TaskLoadTransactions.req:", errReq)
			}
		}

	}

	fmt.Println("TaskLoadTransactions finished", time.Now().Format("2006-01-02 15:04:05"))

	return nil
}
