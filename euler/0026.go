package main

import "fmt"

func main() {
	maxRecursion := -1
	value := -1
	for i := 2; i < 1000; i += 1 {
		recur := recursive(i)
		if recur > maxRecursion {
			maxRecursion = recur
			value = i
		}
	}
	fmt.Println(value, maxRecursion)
}

func recursive(n int) int {
	remain := 1
	digits := make([]int, 0, 1000)
	remains := make([]int, 0, 1000)
	for {
		remain *= 10
		digits = append(digits, remain/n)
		remain = remain % n
		if remain == 0 {
			return 0
		}
		for j := range remains {
			if remains[j] == remain {
				fmt.Println(n, digits, len(remains)-j)
				return len(remains) - j
			}
		}
		remains = append(remains, remain)
	}
}
