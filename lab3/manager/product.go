package manager

type Product struct {
	DistributorID    uint    `gorm:"foreignKey:DistributorID;default:1" json:"distributor_id"`
	Name             string  `gorm:"type:varchar" json:"name"`
	Price            float32 `gorm:"type:numeric(16,2);not null" json:"price"`
	PriceOld         float32 `gorm:"type:numeric(16,2);default:0" json:"price_old"`
	Discount         float32 `gorm:"type:numeric(8,2);default:0" json:"discount"`
	DiscountPeriodID *uint   `gorm:"foreignKey:DiscountPeriodID" json:"discount_period_id"`
	Available        *bool   `gorm:"type:bool;default:true" json:"available"`
	SubCategory      string  `gorm:"foreignKey:SubCategoryID" json:"sub_category"`
	Link             string  `gorm:"type:varchar" json:"link"`
	ImageID          *uint   `gorm:"foreignKey:ImageId" json:"image_id"`
	SpecialCondition string  `gorm:"type:varchar;default:''" json:"special_condition"`
}
