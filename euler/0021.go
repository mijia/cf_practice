package main

import "fmt"

func main() {
	sum := 0
	for i := 2; i < 10000; i += 1 {
		if isAmicableNumber(i) {
			sum += i
			fmt.Println(i)
		}
	}
	fmt.Println(sum)
}

func isAmicableNumber(n int) bool {
	d := divisorSum(n)
	if d == n {
		return false
	}
	return divisorSum(d) == n
}

func divisorSum(n int) int {
	sum := 1
	for i := 2; i <= n/2; i += 1 {
		if n%i == 0 && i < n {
			sum += i
		}
	}
	return sum
}
