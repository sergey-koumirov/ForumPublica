package models

//DvRecord model
type DvRecord struct {
	Description string
	K           float64
}

//Deviation skill model
type Deviation struct {
	ID int32   `gorm:"column:id;primary_key:yes"`
	K  float64 `gorm:"column:k"`
}

//TableName deviation model table name
func (j *Deviation) TableName() string {
	return "fp_deviations"
}
