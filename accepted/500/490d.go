package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
var _ = fmt.Println

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	a1, b1, a2, b2 := ReadInt64(), ReadInt64(), ReadInt64(), ReadInt64()
	a3Count, a2Count := howManyNumbers(a1*b1, 3), howManyNumbers(a1*b1, 2)
	b3Count, b2Count := howManyNumbers(a2*b2, 3), howManyNumbers(a2*b2, 2)

	a3Count, b3Count = subtract(a3Count, b3Count)
	a2Count += a3Count
	b2Count += b3Count
	a2Count, b2Count = subtract(a2Count, b2Count)

	na1, nb1 := bites(a1, b1, a3Count, a2Count)
	na2, nb2 := bites(a2, b2, b3Count, b2Count)
	if na1*nb1 != na2*nb2 {
		PrintInts(-1)
	} else {
		PrintInts64(a2Count + a3Count + b2Count + b3Count)
		writer.WriteByte('\n')
		PrintInts64(na1, nb1)
		writer.WriteByte('\n')
		PrintInts64(na2, nb2)
	}
}

func bites(x, y int64, t3, t2 int64) (int64, int64) {
	for t3 > 0 {
		if x%3 == 0 {
			x = x / 3 * 2
			t3--
		} else if y%3 == 0 {
			y = y / 3 * 2
			t3--
		} else {
			break
		}
	}
	for t2 > 0 {
		if x%2 == 0 {
			x = x / 2
			t2--
		} else if y%2 == 0 {
			y = y / 2
			t2--
		} else {
			break
		}
	}
	return x, y
}

func subtract(x, y int64) (int64, int64) {
	if x > y {
		return x - y, 0
	} else {
		return 0, y - x
	}
}

func howManyNumbers(x, base int64) int64 {
	var n int64 = 0
	for x%base == 0 {
		n++
		x /= base
	}
	return n
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
