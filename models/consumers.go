package models


type Consumers struct {
	Model
	Uuid *string		`json:"uuid"`
	Name *string		`json:"name"`
	Thumb *string		`json:"thumb"`
	Age *uint64			`json:"age"`
}
