package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	var nS string
	fmt.Printf("n = ")
	fmt.Scan(&nS)
	n, _ := new(big.Int).SetString(nS, 10)
	if n.ProbablyPrime(20) {
		fmt.Println("素数です")
	} else {
		fmt.Printf("素因数: ")
		start := time.Now()
		fmt.Println(calcFactor(n))
		end := time.Now()
		fmt.Printf("%fsec\n", (end.Sub(start)).Seconds())
	}
}

func calcFactor(n *big.Int) *big.Int {
	z1 := big.NewInt(2)
	z2 := big.NewInt(2)
	z2_z1 := big.NewInt(0)
	a := big.NewInt(1)
	b := big.NewInt(1)
	result := big.NewInt(1)
	zero := big.NewInt(0)
	one := big.NewInt(1)
	twe := big.NewInt(2)
	three := big.NewInt(3)
	five := big.NewInt(5)
	thrteen := big.NewInt(13)
	module := big.NewInt(0)

	// 2対策
	n.DivMod(n, twe, module)
	if module.Cmp(zero) == 0 {
		n.Mul(n, twe)
		return twe
	}
	n.Mul(n, twe).Add(n, module)

	// 3対策
	n.DivMod(n, three, module)
	if module.Cmp(zero) == 0 {
		n.Mul(n, three)
		return three
	}
	n.Mul(n, three).Add(n, module)

	// 5対策
	n.DivMod(n, five, module)
	if module.Cmp(zero) == 0 {
		n.Mul(n, five)
		return five
	}
	n.Mul(n, five).Add(n, module)

	// 13対策
	n.DivMod(n, thrteen, module)
	if module.Cmp(zero) == 0 {
		n.Mul(n, thrteen)
		return thrteen
	}
	n.Mul(n, thrteen).Add(n, module)

	for {
		z1.Mul(z1, z1)
		z1.Add(z1, one)
		z1.Mod(z1, n)
		z2.Mul(z2, z2)
		z2.Add(z2, one)
		z2.Mod(z2, n)
		z2.Mul(z2, z2)
		z2.Add(z2, one)
		z2.Mod(z2, n)

		result.GCD(a, b, z2_z1.Sub(z2, z1).Abs(z2_z1), n)
		if result.Cmp(one) != 0 {
			break
		}
	}

	return result
}
