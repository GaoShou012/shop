package models

const (
	FranchiseesTableName = "franchisees"
)

type Franchisees struct {
	Model
	Status                *uint64  `json:"status" gorm:"default:0"`
	Username              *string  `json:"username"`
	Password              *string  `json:"password"`
	Balance               *float32 `json:"balance"`
	WithdrawalTotalAmount *float32 `json:"withdrawalTotalAmount"`
	WithdrawalTotalTimes  *uint64  `json:"withdrawalTotalTimes" gorm:"default:0"`
}

func (m *Franchisees) Table() string {
	return "franchisees"
}