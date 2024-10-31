package models

type Distributor struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar" json:"name"`
}

var DistributorSeedData = []Distributor{
	{ID: 1, Name: "Linella"},
	{ID: 2, Name: "Kaufland"},
	{ID: 3, Name: "Local"},
}
