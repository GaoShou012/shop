package models

type ShopsOwners struct {
	Model
	ShopCode *string					`json:"shopCode"`
	FranchiseeName *string				`json:"franchiseeName"`
	PlatformCommissionPercent *uint64	`json:"platformCommissionPercent"`
}
