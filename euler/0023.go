package main

import "fmt"

func main() {
	panic("Not accepted!")
	total := 0
	for n := 25; n <= 28123; n += 1 {
		condition := false
		for i := 12; n-i >= 12; i += 1 {
			if isAbundant(i) && isAbundant(n-i) {
				condition = true
				//fmt.Printf("%d = %d + %d\n", n, i, n-i)
				break
			}
		}
		if !condition {
			total += n
		}
	}
	fmt.Println(total)
}

var cache map[int]bool

func isAbundant(n int) bool {
	if is, ok := cache[n]; ok {
		return is
	}
	is := n < divisorSum(n)
	cache[n] = is
	return is
}

func divisorSum(n int) int {
	sum := 1
	for i := 2; i <= n/2; i += 1 {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func init() {
	cache = make(map[int]bool)
}
