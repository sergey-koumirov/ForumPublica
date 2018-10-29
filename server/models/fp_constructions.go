package models

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/db"
	"encoding/json"
)

type Construction struct {
	Id     int64 `gorm:"column:id;primary_key"`
	UserId int64 `gorm:"column:user_id"`

	Name        string `gorm:"column:name"`
	CitadelType string `gorm:"column:citadel_type"`
	RigFactor   string `gorm:"column:rig_factor"`
	SpaceType   string `gorm:"column:space_type"`

	Bpos ConstructionBpos    `gorm:"foreignkey:ConstructionId"`
	Runs ConstructionBpoRuns `gorm:"foreignkey:ConstructionId"`
}

func (u *Construction) TableName() string {
	return "fp_constructions"
}

func (m *Construction) Delete() {
	for _, bp := range m.Bpos {
		bp.Delete()
	}
	db.DB.Delete(&m)
}

type ConstructionBpo struct {
	Id             int64  `gorm:"column:id;primary_key"`
	ConstructionId int64  `gorm:"column:construction_id"`
	TransactionId  int64  `gorm:"column:transaction_id"`
	Kind           string `gorm:"column:kind"`
	TypeId         int32  `gorm:"column:type_id"`
	ME             int32  `gorm:"column:me"`
	TE             int32  `gorm:"column:te"`
	Qty            int64  `gorm:"column:qty"`
}

func (m *ConstructionBpo) TableName() string {
	return "fp_construction_bpos"
}

func (m *ConstructionBpo) TypeName() string {
	return static.Types[m.TypeId].Name
}

func (u *ConstructionBpo) MarshalJSON() ([]byte, error) {
	type Alias ConstructionBpo
	return json.Marshal(&struct {
		TypeName string `json:"TypeName"`
		*Alias
	}{
		TypeName: u.TypeName(),
		Alias:    (*Alias)(u),
	})
}

func (m *ConstructionBpo) Delete() {
	db.DB.Delete(&m)
}

type ConstructionBpos []ConstructionBpo

type ConstructionBpoRun struct {
	Id                int64  `gorm:"column:id;primary_key"`
	ConstructionId    int64  `gorm:"column:construction_id"`
	TypeId            int32  `gorm:"column:type_id"`
	ConstructionBpoId int64  `gorm:"column:construction_bpo_id"`
	Repeats           int32  `gorm:"column:repeats"`
	ME                int32  `gorm:"column:me"`
	TE                int32  `gorm:"column:te"`
	Qty               int64  `gorm:"column:qty"`
	CitadelType       string `gorm:"column:citadel_type"`
	RigFactor         string `gorm:"column:rig_factor"`
	SpaceType         string `gorm:"column:space_type"`
}

type ConstructionBpoRuns []ConstructionBpoRun

func (u *ConstructionBpoRun) TableName() string {
	return "fp_construction_bpo_runs"
}

type ConstructionExpense struct {
	Id                int64   `gorm:"column:id;primary_key"`
	ConstructionBpoId int64   `gorm:"column:construction_bpo_id"`
	Description       string  `gorm:"column:description"`
	ExValue           float64 `gorm:"column:exvalue"`
}

func (u *ConstructionExpense) TableName() string {
	return "fp_construction_expenses"
}
