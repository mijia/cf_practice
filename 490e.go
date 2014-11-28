package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
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
		pattern := ReadString()
		if num, ok := getNumber(pattern, prevMax); !ok {
			writer.WriteString("NO")
			return
		} else {
			outputs[i] = num
			prevMax = num
		}
	}
	writer.WriteString("YES\n")
	for _, n := range outputs {
		PrintInts(n)
		writer.WriteString("\n")
	}

}

func getNumber(pattern string, prevMax int) (int, bool) {
	if num, err := strconv.Atoi(pattern); err == nil {
		return num, num > prevMax
	}
	// we found some ? mark
	low, high := getRange(pattern)
	fmt.Println(low, high, pattern)
	if low >= high || high <= prevMax {
		return -1, false
	}
	if low > prevMax {
		return low, true
	}
	return 1, true
}

func getRange(pattern string) (int, int) {
	low, high := 0, 0
	if pattern[0] == '?' {
		low, high = 1, 9
	} else {
		val, _ := strconv.Atoi(string(pattern[0]))
		low, high = val, val
	}
	for i := 1; i < len(pattern); i += 1 {
		if pattern[i] == '?' {
			low = low * 10
			high = high * 10 + 9
		} else {
			val, _ := strconv.Atoi(string(pattern[i]))
			low = low * 10 + val
			high = high * 10 + val
		}
	}
	return low, high
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
	for _, value := range ints {
		writer.WriteString(strconv.FormatInt(value, 10))
		writer.WriteByte(' ')
	}
}

func PrintInts(ints ...int) {
	for _, value := range ints {
		writer.WriteString(strconv.Itoa(value))
		writer.WriteByte(' ')
	}
}
