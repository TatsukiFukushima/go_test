package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		fib(100)
	}
	end := time.Now()
	fmt.Printf("%fsec\n", (end.Sub(start)).Seconds())
}

func fib(n int) uint64 {
	result := make([]uint64, n)

	result[0] = 1
	if n >= 2 {
		result[1] = 1
	}

	for i := 2; i < n; i++ {
		result[i] = result[i-1] + result[i-2]
	}

	return result[n-1]
}

func fibAppend(n int) uint64 {
	var result []uint64

	result = append(result, 1)
	if n >= 2 {
		result = append(result, 1)
	}

	for i := 2; i < n; i++ {
		result = append(result, result[i-1]+result[i-2])
	}

	return result[n-1]
}
