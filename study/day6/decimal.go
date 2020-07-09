package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
)

func main() {
	aa := 33.11
	res := decimal.NewFromFloat(aa).Mul(decimal.NewFromInt(2))
	fmt.Println(res.Float64())

	bb := 33.99
	fmt.Println(math.Round(bb)) // 输出34
	fmt.Println(math.Round(bb*100)/100) // 输出 33.99
	fmt.Println(math.Floor(bb))  // 输出33
}
