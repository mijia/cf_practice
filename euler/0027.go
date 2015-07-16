package main

import "fmt"

func main() {
	maxN, maxA, maxB := 0, 0, 0
	for b := 1; b < 1000; b += 1 {
		for a := -999; a < 1000; a += 1 {
			posibleN := maxPrimeN(a, b)
			if posibleN > maxN {
				maxN = posibleN
				maxA = a
				maxB = b
				fmt.Println(maxN, maxA, maxB)
			}
		}
	}
	fmt.Println(maxA * maxB)
}

func maxPrimeN(a, b int) int {
	for n := 0; ; n += 1 {
		pn := n*n + a*n + b
		if !isPrime(pn) {
			return n - 1
		}
	}
}

var cache map[int]bool

func isPrime(n int) bool {
	if n < 0 {
		return false
	}
	if n == 1 || n == 2 || n == 3 {
		return true
	}

	if is, ok := cache[n]; ok {
		return is
	}

	for i := 2; i*i < n; i += 1 {
		if n%i == 0 {
			cache[n] = false
			return false
		}
	}
	cache[n] = true
	return true
}

func init() {
	cache = make(map[int]bool)
}
