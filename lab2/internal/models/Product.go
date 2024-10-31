package models

import "time"

type Product struct {
	ID               uint    `gorm:"primaryKey" json:"id"`
	DistributorID    uint    `gorm:"foreignKey:DistributorID" json:"distributor_id"`
	Name             string  `gorm:"type:varchar" json:"name"`
	Price            float32 `gorm:"type:numeric(16,2)" json:"price"`
	PriceOld         float32 `gorm:"type:numeric(16,2);default:0" json:"price_old"`
	Discount         float32 `gorm:"type:numeric(8,2);default:0" json:"discount"`
	DiscountPeriodID *uint   `gorm:"foreignKey:DiscountPeriodID" json:"discount_period_id"`
	Available        *bool   `gorm:"default:true" json:"available"`
	SubCategoryID    uint    `gorm:"foreignKey:SubCategoryID" json:"sub_category_id"`
	Link             string  `gorm:"type:varchar" json:"link"`
	Image            *[]byte `gorm:"type:bytea" json:"image"`
	SpecialCondition string  `gorm:"type:varchar;default:''" json:"special_condition"`
}

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar" json:"name"`
}

type SubCategory struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	CategoryID uint   `gorm:"foreignKey:CategoryID" json:"category_id"`
}

type DiscountPeriod struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type Distributor struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar" json:"name"`
}
