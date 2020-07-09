package models

import "time"

const (
	UsersOrdersTable = "users_orders"
)

type UsersOrders struct {
	Model
	PlatformOrderNo  *uint64    `json:"platformOrderNo"`
	UserId           *uint64    `json:"consumerId"`
	UserThumb        *string    `json:"userThumb"`
	Status           *uint64    `json:"status" gorm:"default:0"`
	TotalImportPrice *float64   `json:"totalImportPrice"`
	TotalExportPrice *float64   `json:"totalExportPrice"`
	ShopCode         *string    `json:"shopCode"`
	ShopAddr         *string    `json:"shopAddr"`
	ClientToken      *string    `json:"clientToken"`
	PaidAt           *time.Time `json:"paidAt" gorm:"default:NULL"`
}
