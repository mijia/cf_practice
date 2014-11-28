package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	pattern := ReadString()
	prevMax, _, _, _ := getPatternMeta(pattern)
	outputs[0] = prevMax
	for i := 1; i < N; i += 1 {
		pattern = ReadString()
		if num, ok := getNumber(pattern, prevMax); !ok {
			writer.WriteString("NO")
			return
		} else {
			outputs[i] = num
			prevMax = num
		}
	}
	writer.WriteString("YES")
	for _, n := range outputs {
		writer.WriteString("\n")
		PrintInts(n)
	}
}

func getNumber(pattern string, prevMax int) (int, bool) {
	if num, err := strconv.Atoi(pattern); err == nil {
		return num, num > prevMax
	}
	// we found some ? mark
	low, high, start, end := getPatternMeta(pattern)
	if low >= high || high <= prevMax {
		return -1, false
	}
	if low > prevMax {
		return low, true
	}
	return findNumber(pattern, prevMax, start, end)
}

func adjustStart(pattern string, prevMax int) int {
	digits := make([]string, 0)
	for i := len(pattern) - 1; i >= 0; i -= 1 {
		mod := prevMax % 10
		prevMax /= 10
		if pattern[i] == '?' {
			digits = append([]string{strconv.Itoa(mod)}, digits...)
		}
	}
	v, _ := strconv.Atoi(strings.Join(digits, ""))
	return v
}

func findNumber(pattern string, prevMax, low, high int) (int, bool) {
	start := low
	end := high
	for start <= end {
		mid := start + (end-start)/2
		midN := createNumber(pattern, mid)
		if midN == prevMax {
			start = mid + 1
			break
		}
		if midN > prevMax {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if start > high {
		return -1, false
	}
	num := createNumber(pattern, start)
	return num, num > prevMax
}

func createNumber(pattern string, mask int) int {
	newData := []byte(pattern)
	for i := len(newData) - 1; i >= 0; i -= 1 {
		if newData[i] == '?' {
			mod := mask % 10
			mask /= 10
			if mod == 0 && i == 0 {
				newData[i] = byte(1)
			} else {
				newData[i] = byte(mod)
			}
		} else {
			newData[i] -= '0'
		}
	}
	ans := 0
	for i := 0; i < len(newData); i += 1 {
		ans = ans * 10 + int(newData[i])
	}
	return ans
}

func getPatternMeta(pattern string) (low, high, start, end int) {
	qDigits := 0
	if pattern[0] == '?' {
		low, high = 1, 9
		start = 1
		qDigits++
	} else {
		val, _ := strconv.Atoi(string(pattern[0]))
		low, high = val, val
	}
	for i := 1; i < len(pattern); i += 1 {
		if pattern[i] == '?' {
			low = low * 10
			high = high*10 + 9
			qDigits++
		} else {
			val, _ := strconv.Atoi(string(pattern[i]))
			low = low*10 + val
			high = high*10 + val
		}
	}
	end = 9
	for i := 1; i < qDigits; i += 1 {
		start *= 10
		end = end*10 + 9
	}

	return
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
