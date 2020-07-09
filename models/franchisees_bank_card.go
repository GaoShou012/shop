package models

const (
	FranchiseesBankCardTable = "franchisees_bank_card"
)

type FranchiseesBankCard struct {
	Model
	Username   *string
	RealName   *string
	BankName   *string
	BankBranch *string
	CardNo     *string
	Phone      *string
}
