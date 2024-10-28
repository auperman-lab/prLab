package models

type Product struct {
	ID               uint    `gorm:"primaryKey"`                 // PK
	DistributorID    uint    `gorm:"not null"`                   // FK to Distributor
	Name             string  `gorm:"type:varchar(255);not null"` // Product name
	Price            float64 `gorm:"type:real"`                  // Product price
	PriceOld         float64 `gorm:"type:real"`                  // Old price
	Discount         float64 `gorm:"type:real"`                  // Discount amount
	DiscountID       uint    `gorm:"not null"`                   // FK to Discount
	Available        bool    `gorm:"default:true"`               // Availability status
	SubCategoryID    uint    `gorm:"not null"`                   // FK to SubCategory
	Link             string  `gorm:"type:varchar(255)"`          // Product link
	Image            []byte  `gorm:"type:bytea"`                 // Image stored as byte array
	SpecialCondition string  `gorm:"type:varchar(255)"`          // Special condition
}
