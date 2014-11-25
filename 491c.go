// http://en.wikipedia.org/w/index.php?title=Hungarian_algorithm

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var _ = fmt.Println
var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

type MapSorter struct {
	keys []string
	data map[string]int
}

func (s *MapSorter) Len() int           { return len(s.keys) }
func (s *MapSorter) Swap(x, y int)      { s.keys[x], s.keys[y] = s.keys[y], s.keys[x] }
func (s *MapSorter) Less(x, y int) bool { return s.data[s.keys[x]] < s.data[s.keys[y]] }

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	_, K := ReadInt32(), ReadInt32()
	kCodes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"[0:K]
	cipherText, correctText := ReadString(), ReadString()

	counts := make(map[string]int)
	for i := range cipherText {
		ch := string(cipherText[i])
		cch := string(correctText[i])
		counts[ch+cch] += 1
	}

	sorter := &MapSorter{data: counts}
	sorter.keys = make([]string, len(counts))
	index := 0
	for k := range counts {
		sorter.keys[index] = k
		index++
	}
	sort.Sort(sort.Reverse(sorter))
	cipher := make(map[string]string)
	targets := make(map[string]bool)
	for _, k := range sorter.keys {
		ch := string(k[0])
		target := string(k[1])
		if _, ok := cipher[ch]; ok {
			continue
		}
		if _, ok := targets[target]; ok {
			continue
		}
		cipher[ch] = target
		targets[target] = true
	}

	correctCount := 0
	for c, t := range cipher {
		correctCount += counts[c+t]
	}
	PrintInts(correctCount)
	writer.WriteString("\n")
	for _, kC := range kCodes {
		writer.WriteString(cipher[string(kC)])
	}
	writer.WriteString("\n")
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

func PrintInts(ints ...int) {
	for _, value := range ints {
		writer.WriteString(strconv.Itoa(value))
		writer.WriteByte(' ')
	}
}

func PrintStrings(strings ...string) {
	for _, value := range strings {
		writer.WriteString(value)
		writer.WriteByte(' ')
	}
}
