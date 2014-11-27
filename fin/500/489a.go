package main

import (
	"bufio"
	"os"
	"strconv"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n := ReadInt32()
	data := make([]int, n)
	for i := 0; i < n; i += 1 {
		data[i] = ReadInt32()
	}
	swaps := 0
	swapIndex := make([][]int, 0)
	for i := 0; i < n-1; i += 1 {
		minValue := 1000000001
		minIndex := -1
		for j := i + 1; j < n; j += 1 {
			if data[j] < minValue {
				minValue = data[j]
				minIndex = j
			}
		}
		if minValue < data[i] {
			swaps += 1
			data[i], data[minIndex] = data[minIndex], data[i]
			swapIndex = append(swapIndex, []int{i, minIndex})
		}
	}
	PrintInts(swaps)
	writer.WriteByte('\n')
	for _, indexs := range swapIndex {
		PrintInts(indexs...)
		writer.WriteByte('\n')
	}
}

func ReadInt32() int {
	scanner.Scan()
	ans, _ := strconv.Atoi(scanner.Text())
	return ans
}

func PrintInts(ints ...int) {
	for _, value := range ints {
		writer.WriteString(strconv.Itoa(value))
		writer.WriteByte(' ')
	}
}
