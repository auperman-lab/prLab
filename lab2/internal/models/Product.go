package models

import "time"

type Product struct {
	ID               uint    `gorm:"primaryKey" json:"id"`
	DistributorID    uint    `gorm:"foreignKey:DistributorID;default:1" json:"distributor_id"`
	Name             string  `gorm:"type:varchar" json:"name"`
	Price            float32 `gorm:"type:numeric(16,2);not null" json:"price"`
	PriceOld         float32 `gorm:"type:numeric(16,2);default:0" json:"price_old"`
	Discount         float32 `gorm:"type:numeric(8,2);default:0" json:"discount"`
	DiscountPeriodID *uint   `gorm:"foreignKey:DiscountPeriodID" json:"discount_period_id"`
	Available        *bool   `gorm:"type:bool;default:true" json:"available"`
	SubCategoryID    uint    `gorm:"foreignKey:SubCategoryID" json:"sub_category_id"`
	Link             string  `gorm:"type:varchar" json:"link"`
	ImageID          *uint   `gorm:"foreignKey:ImageId" json:"image_id"`
	SpecialCondition string  `gorm:"type:varchar;default:''" json:"special_condition"`
}

type DiscountPeriod struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type Image struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Image []byte `gorm:"type:bytea" json:"image"`
}

type ReturnProduct struct {
	ID                  uint      `json:"id"`
	Distributor         string    `json:"distributor"`
	Name                string    `json:"name"`
	Price               float32   `json:"price"`
	PriceOld            float32   `json:"price_old"`
	Discount            float32   `json:"discount"`
	DiscountPeriodStart time.Time `json:"discount_period_start"`
	DiscountPeriodEnd   time.Time `json:"discount_period_end"`
	Available           *bool     `json:"available"`
	Category            string    `json:"category"`
	SubCategory         string    `json:"sub_category"`
	Link                string    `json:"link"`
	SpecialCondition    string    `json:"special_condition"`
}
