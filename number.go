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
	one := big.NewInt(1)
	numbers := []*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(5), big.NewInt(7), big.NewInt(11), big.NewInt(13)}

	// 小さい値が素因数だとたまにエラーが起こるので対策
	for _, number := range numbers {
		if isModZero(n, number) {
			return number
		}
	}

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

// isModZero 余りがゼロかどうかを判定
func isModZero(n, m *big.Int) bool {
	zero := big.NewInt(0)
	module := big.NewInt(0)
	n.DivMod(n, m, module)
	if module.Cmp(zero) == 0 {
		n.Mul(n, m)
		return true
	}
	n.Mul(n, m).Add(n, module)
	return false
}
