package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	splited_s := strings.Fields(s)
	result := make(map[string]int)
	for _, e := range splited_s {
		if _, ok := result[e]; !ok {
			result[e] = 0
		}
		result[e] += 1
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
