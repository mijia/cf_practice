package main

import (
	"fmt"
	"strings"
)

func main() {
	sum := 0
	for i := 1; i <= 1000; i += 1 {
		sum += len(NumberToBritishWords(i))
	}
	fmt.Println(sum)
}

var digits map[int]string
var digits2 map[int]string

func NumberToBritishWords(n int) string {
	above100, below100 := n/100, n%100
	parts := make([]string, 0)
	if below100 < 20 {
		parts = append(parts, digits[below100])
	} else {
		d2, d1 := below100/10, below100%10
		parts = append(parts, digits[d1])
		parts = append(parts, digits2[d2])
	}
	if above100 > 0 {
		if below100 > 0 {
			parts = append(parts, "and")
		}
		if above100 >= 10 {
			parts = append(parts, "onethousand")
		} else {
			parts = append(parts, digits[above100], "hundred")
		}
	}
	return strings.Join(parts, "")
}

func init() {
	digits = make(map[int]string)
	digits[0] = ""
	digits[1] = "one"
	digits[2] = "two"
	digits[3] = "three"
	digits[4] = "four"
	digits[5] = "five"
	digits[6] = "six"
	digits[7] = "seven"
	digits[8] = "eight"
	digits[9] = "nine"
	digits[10] = "ten"
	digits[11] = "eleven"
	digits[12] = "twelve"
	digits[13] = "thirteen"
	digits[14] = "fourteen"
	digits[15] = "fifteen"
	digits[16] = "sixteen"
	digits[17] = "seventeen"
	digits[18] = "eighteen"
	digits[19] = "nineteen"

	digits2 = make(map[int]string)
	digits2[2] = "twenty"
	digits2[3] = "thirty"
	digits2[4] = "forty"
	digits2[5] = "fifty"
	digits2[6] = "sixty"
	digits2[7] = "seventy"
	digits2[8] = "eighty"
	digits2[9] = "ninety"
}
