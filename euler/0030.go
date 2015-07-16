package main

import "fmt"

func main() {
	total := 0
	base := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	for i := 2; i <= 999999; i += 1 {
		sum := 0
		for j := 1; j <= len(base)-1; j += 1 {
			if i*10 < base[j] {
				break
			}
			digit := i % base[j] / base[j-1]
			sum += digit * digit * digit * digit * digit
			if sum > i {
				break
			}
		}
		if sum == i {
			fmt.Println(i)
			total += sum
		}
	}
	fmt.Println(total)
}
