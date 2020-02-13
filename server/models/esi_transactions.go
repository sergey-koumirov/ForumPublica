package models

//Transaction transaction model
type Transaction struct {
	ID             int64   `gorm:"column:id;primary_key"`
	EsiCharacterID int64   `gorm:"column:esi_character_id"`
	ClientID       int64   `gorm:"column:client_id"`
	Dt             string  `gorm:"column:dt"`
	IsBuy          bool    `gorm:"column:is_buy"`
	IsPersonal     bool    `gorm:"column:is_personal"`
	JournalRefID   int64   `gorm:"column:journal_ref_id"`
	LocationID     int64   `gorm:"column:location_id"`
	Quantity       int64   `gorm:"column:quantity"`
	TypeID         int32   `gorm:"column:type_id"`
	UnitPrice      float64 `gorm:"column:unit_price"`

	Character  *Character  `gorm:"foreignkey:EsiCharacterID"`
	ClientName *ClientName `gorm:"foreignkey:ClientID"`
	Location   *Location   `gorm:"foreignkey:LocationID"`
}

//TableName transaction model table name
func (j *Transaction) TableName() string {
	return "esi_transactions"
}
