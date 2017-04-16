package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	moji := strings.Fields(s)
	num := len(moji)

	m := map[string]int{}

	for i := 0; i < num; i++ {
		m[moji[i]]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
