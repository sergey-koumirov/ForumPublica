package models

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"encoding/json"
)

//Construction model
type Construction struct {
	ID     int64 `gorm:"column:id;primary_key"`
	UserID int64 `gorm:"column:user_id"`

	Name        string `gorm:"column:name"`
	CitadelType string `gorm:"column:citadel_type"`
	RigFactor   string `gorm:"column:rig_factor"`
	SpaceType   string `gorm:"column:space_type"`

	Bpos ConstructionBpos    `gorm:"foreignkey:ConstructionId"`
	Runs ConstructionBpoRuns `gorm:"foreignkey:ConstructionId"`
}

//TableName construction table name
func (m *Construction) TableName() string {
	return "fp_constructions"
}

//Delete delete model and children
func (m *Construction) Delete() {
	for _, bp := range m.Bpos {
		bp.Delete()
	}
	for _, r := range m.Runs {
		r.Delete()
	}
	db.DB.Delete(&m)
}

//ConstructionBpo construction BPO model
type ConstructionBpo struct {
	ID            int64  `gorm:"column:id;primary_key"`
	TransactionID int64  `gorm:"column:transaction_id"`
	Kind          string `gorm:"column:kind"`
	TypeID        int32  `gorm:"column:type_id"`
	ME            int32  `gorm:"column:me"`
	TE            int32  `gorm:"column:te"`
	Qty           int64  `gorm:"column:qty"`

	ConstructionID int64 `gorm:"column:construction_id"`
	Construction   *Construction

	Expenses ConstructionExpenses `gorm:"foreignkey:ConstructionBpoId"`
}

//TableName construction BPO table name
func (m *ConstructionBpo) TableName() string {
	return "fp_construction_bpos"
}

//TypeName construction BPO type name
func (m *ConstructionBpo) TypeName() string {
	return static.Types[m.TypeID].Name
}

//MarshalJSON custom marshalling
func (m *ConstructionBpo) MarshalJSON() ([]byte, error) {
	type Alias ConstructionBpo
	return json.Marshal(&struct {
		TypeName string `json:"TypeName"`
		*Alias
	}{
		TypeName: m.TypeName(),
		Alias:    (*Alias)(m),
	})
}

//Delete delete model and children
func (m *ConstructionBpo) Delete() {
	for _, e := range m.Expenses {
		e.Delete()
	}
	db.DB.Delete(&m)
}

//ConstructionBpos array
type ConstructionBpos []ConstructionBpo

//ConstructionBpoRun construction BPO run model
type ConstructionBpoRun struct {
	ID                int64  `gorm:"column:id;primary_key"`
	ConstructionID    int64  `gorm:"column:construction_id"`
	TypeID            int32  `gorm:"column:type_id"`
	ConstructionBpoID int64  `gorm:"column:construction_bpo_id"`
	ME                int32  `gorm:"column:me"`
	TE                int32  `gorm:"column:te"`
	Repeats           int32  `gorm:"column:repeats"`
	Qty               int64  `gorm:"column:qty"`
	ExactQty          int64  `gorm:"-"`
	CitadelType       string `gorm:"column:citadel_type"`
	RigFactor         string `gorm:"column:rig_factor"`
	SpaceType         string `gorm:"column:space_type"`
}

//ConstructionBpoRuns array
type ConstructionBpoRuns []ConstructionBpoRun

//TableName table name
func (u *ConstructionBpoRun) TableName() string {
	return "fp_construction_bpo_runs"
}

//Total run total
func (u *ConstructionBpoRun) Total() int64 {
	return int64(u.Repeats) * u.Qty
}

//Delete delete model
func (u *ConstructionBpoRun) Delete() {
	db.DB.Delete(&u)
}

//ConstructionExpense table
type ConstructionExpense struct {
	ID          int64   `gorm:"column:id;primary_key"`
	Description string  `gorm:"column:description"`
	ExValue     float64 `gorm:"column:exvalue"`

	ConstructionBpoID int64 `gorm:"column:construction_bpo_id"`
	ConstructionBpo   *ConstructionBpo
}

//TableName for ConstructionExpense table
func (m *ConstructionExpense) TableName() string {
	return "fp_construction_expenses"
}

//ConstructionExpenses array
type ConstructionExpenses []ConstructionExpense

//Delete delete model
func (m *ConstructionExpense) Delete() {
	db.DB.Delete(&m)
}
