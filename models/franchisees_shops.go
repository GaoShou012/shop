package models

const (
	FranchiseesShopsTable = "franchisees_shops"
)

type FranchiseesShops struct {
	Model
	Enable                    *bool   `json:"enable" gorm:"default:false"`
	FranchiseeUsername        *string `json:"franchiseeUsername"`
	ShopCode                  *string `json:"shopCode"`
	PlatformCommissionPercent *uint64 `json:"platformCommissionPercent"`
}
