package services

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"ForumPublica/server/models"
	"fmt"
	"time"
)

var sqlByDate = `
select d.date as d, sum(t.quantity * t.unit_price) as v
  from sys_dates d
         left join esi_transactions t on substring(t.dt,1,10) = d.date 
                                     and t.is_buy=0
									 and t.esi_character_id in (
										  select c.id 
											from esi_characters c
											where c.user_id = ? 
									 )
  where d.date >= ?
  group by d.date
  order by d.date`

var sqlByType = `
select t.type_id, sum(t.quantity) as q, sum(t.quantity * t.unit_price) as v
  from esi_transactions t
  where t.dt >= ?
    and t.is_buy=0
    and t.esi_character_id in (
      select c.id 
        from esi_characters c
        where c.user_id = ? 
    )
  group by t.type_id
  order by sum(t.quantity * t.unit_price) desc`

//TransactionsSummary list
func TransactionsSummary(userID int64) models.TrSummary {

	minus30d := time.Now().AddDate(0, 0, -30).Format("2006-01-02")

	rowsByType, errByType := db.DB.Raw(sqlByType, minus30d, userID).Rows()
	defer rowsByType.Close()
	if errByType != nil {
		fmt.Println("loadMarketData.errByType:", errByType)
	}

	recordsByType := make([]models.TrByType, 0)
	total := float64(0)
	for rowsByType.Next() {
		temp := models.TrByType{}
		rowsByType.Scan(&temp.TypeID, &temp.TotalQty, &temp.TotalValue)
		temp.TypeName = static.Types[temp.TypeID].Name
		recordsByType = append(recordsByType, temp)
		total = total + temp.TotalValue
	}

	rowsByDate, errByDate := db.DB.Raw(sqlByDate, userID, minus30d).Rows()
	defer rowsByDate.Close()
	if errByDate != nil {
		fmt.Println("loadMarketData.errByDate:", errByDate)
	}

	recordsByDate := make([]models.TrByDate, 0)
	for rowsByDate.Next() {
		temp := models.TrByDate{}
		rowsByDate.Scan(&temp.Dt, &temp.TotalValue)
		recordsByDate = append(recordsByDate, temp)
	}

	result := models.TrSummary{
		Total:  total,
		ByDate: recordsByDate,
		ByType: recordsByType,
	}

	return result
}
