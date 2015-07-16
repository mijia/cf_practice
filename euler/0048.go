package main

import (
	"fmt"
	"math/big"
)

func main() {
	remain := 0
	for i := 1; i <= 1000; i += 1 {
		remain += selfPowerRemain(i)
		remain = remain % 1e10
	}
	fmt.Println(remain)
}

func selfPowerRemain(n int) int {
	mul := 1
	for i := 0; i < n; i += 1 {
		mul *= n
		mul = mul % 1e10
	}
	return mul
}

func bigNumber() {
	sum := big.NewInt(0)
	divider := selfPower(10)
	for i := 1; i <= 1000; i += 1 {
		sp := selfPower(i)
		remain := new(big.Int)
		remain.Mod(sp, divider)
		sum.Add(sum, remain)
	}
	sum.Mod(sum, divider)
	fmt.Println(sum)
}

func selfPower(n int) *big.Int {
	x := big.NewInt(int64(n))
	mul := big.NewInt(1)
	for i := 0; i < n; i += 1 {
		mul.Mul(mul, x)
	}
	return mul
}
