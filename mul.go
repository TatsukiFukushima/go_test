package main

import (
	"fmt"
	"math/big"
)

func main() {
	var xS string
	var yS string
	fmt.Printf("x = ")
	fmt.Scan(&xS)
	fmt.Printf("y = ")
	fmt.Scan(&yS)
	x, _ := new(big.Int).SetString(xS, 10)
	y, _ := new(big.Int).SetString(yS, 10)
	fmt.Println(x.Mul(x, y))
}
