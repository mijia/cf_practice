package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("data/p022_names.txt")
	if err != nil {
		panic(err)
	}

	words := strings.Split(string(data), ",")
	for i := range words {
		words[i] = words[i][1 : len(words[i])-1]
	}
	sort.Strings(words)

	sum := 0
	for i, word := range words {
		sum += nameScore(i+1, word)
	}
	fmt.Println(sum)
}

func nameScore(index int, word string) int {
	data := []byte(word)
	score := 0
	for _, ch := range data {
		score += int(ch - 64)
	}
	return index * score
}
