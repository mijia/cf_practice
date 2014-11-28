package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var _ = fmt.Println
var _ = sort.Sort

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	N := ReadInt32()
	outputs := make([]int, N)
	prevMax := 0
	for i := 0; i < N; i += 1 {
		pattern := []byte(ReadString())
		if num := getPatternMax(pattern); num <= prevMax {
			writer.WriteString("NO")
			return
		} else {
			for j := 0; j < len(pattern); j += 1 {
				if pattern[j] == '?' {
					for k := 0; k < 10; k += 1 {
						if j == 0 && k == 0 {
							continue
						}
						pattern[j] = byte(k) + '0'
						if getPatternMax(pattern) > prevMax {
							break
						}
					}
				}
			}
			n := getPatternMax(pattern)
			prevMax = n
			outputs[i] = n
		}
	}
	writer.WriteString("YES")
	for _, n := range outputs {
		writer.WriteString("\n")
		PrintInts(n)
	}
}

func getPatternMax(pattern []byte) int {
	n := 0
	for i := 0; i < len(pattern); i += 1 {
		n *= 10
		if pattern[i] == '?' {
			n += 9
		} else {
			n += int(pattern[i] - '0')
		}
	}
	return n
}

func ReadString() string {
	scanner.Scan()
	return scanner.Text()
}

func ReadInt32() int {
	scanner.Scan()
	ans, _ := strconv.Atoi(scanner.Text())
	return ans
}

func ReadInt64() int64 {
	scanner.Scan()
	ans, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return ans
}

func PrintInts64(ints ...int64) {
	for i, value := range ints {
		writer.WriteString(strconv.FormatInt(value, 10))
		if i != len(ints)-1 {
			writer.WriteByte(' ')
		}
	}
}

func PrintInts(ints ...int) {
	for i, value := range ints {
		writer.WriteString(strconv.Itoa(value))
		if i != len(ints)-1 {
			writer.WriteByte(' ')
		}
	}
}
