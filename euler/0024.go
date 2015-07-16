package main

import "fmt"

func main() {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	selection := []int{}
	leftRanks := 1000000
	for leftRanks > 0 {
		posCounts := posibleCounts(len(digits) - 1)
		index := leftRanks / posCounts
		remain := leftRanks % posCounts
		if remain == 0 {
			index -= 1
		}
		selection = append(selection, digits[index])
		newDigits := make([]int, 0, len(digits)-1)
		for i, d := range digits {
			if i != index {
				newDigits = append(newDigits, d)
			}
		}
		digits = newDigits
		fmt.Println(leftRanks, posCounts, index, digits)
		if remain == 0 {
			// we can return the sorted results
			for i := len(digits) - 1; i >= 0; i -= 1 {
				selection = append(selection, digits[i])
			}
			break
		}
		leftRanks -= posCounts * index
	}
	fmt.Println(selection)
}

func posibleCounts(n int) int {
	total := 1
	for i := n; i > 1; i -= 1 {
		total *= i
	}
	return total
}
