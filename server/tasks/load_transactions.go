package tasks

import (
	"ForumPublica/server/db"
	"ForumPublica/server/esi"
	"ForumPublica/server/models"
	"ForumPublica/server/services"
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
		fmt.Println(char.ID, char.Name)
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
				created = processBatch(r, char, api)
				fmt.Printf("  from_id: %d added %d\n", from_id, created)
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

func processBatch(r *esi.CharactersWalletTransactionsResponse, char models.Character, api esi.ESI) int {
	trIds := make([]int64, 0)
	for _, t := range r.R {
		trIds = append(trIds, t.TransactionId)
	}

	exTrIds := make([]int64, 0)
	db.DB.Model(&models.Transaction{}).Where("esi_character_id = ? and id in (?)", char.ID, trIds).Pluck("ID", &exTrIds)

	uniqClientIdsMap := make(map[int64]int32)
	uniqLocationIdsMap := make(map[int64]int32)

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
			uniqLocationIdsMap[t.LocationId] = 1
		}
	}

	uniqClientIds := make([]int64, 0)
	for k, _ := range uniqClientIdsMap {
		uniqClientIds = append(uniqClientIds, k)
	}

	exClIds := make([]int64, 0)
	db.DB.Model(&models.ClientName{}).Where("id in (?)", uniqClientIds).Pluck("ID", &exClIds)

	notExClIds := utils.DiffInt64(uniqClientIds, exClIds)

	for _, batch := range utils.InBatchesInt64(notExClIds, 500) {
		names, errN := api.UniverseNames(batch)
		if errN != nil {
			fmt.Println("TaskLoadTransactions.errN:", errN)
		} else {
			for _, name := range names {
				temp := models.ClientName{
					ID:   name.ID,
					Name: name.Name,
				}
				errCr := db.DB.Create(&temp)
				if errCr != nil {
					fmt.Println("TaskLoadTransactions.errCr:", errCr)
				}
			}
		}
	}

	uniqLocationIds := make([]int64, 0)
	for k, _ := range uniqLocationIdsMap {
		uniqLocationIds = append(uniqLocationIds, k)
	}

	exLocIds := make([]int64, 0)
	db.DB.Model(&models.Location{}).Where("id in (?)", uniqLocationIds).Pluck("ID", &exLocIds)

	notExLocIds := utils.DiffInt64(uniqLocationIds, exLocIds)

	for _, lid := range notExLocIds {
		services.AddLocation(api, lid, "", 0, 0)
		fmt.Println("    new location:", lid)
	}

	return len(trIds) - len(exTrIds)
}
