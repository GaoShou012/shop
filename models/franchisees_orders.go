package models

type FranchiseesOrders struct {
	Model
	Status                    *uint64	`gorm:"default:0"`
	ShopCode                  *string
	UserId                    *uint64
	UserThumb                 *string
	TotalImportPrice          *float64
	TotalExportPrice          *float64
	FranchiseeUsername        *string
	FranchiseeCommission      *float64
	PlatformOrderNo           *uint64
	PlatformCommissionPercent *uint64
	PlatformCommission        *float64
}
