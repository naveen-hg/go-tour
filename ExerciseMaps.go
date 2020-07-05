package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	sl := strings.Split(s, " ")
	ret := make(map[string]int)
	for _, v := range(sl) {
		elem, ok := ret[v]
		if ok {
			ret[v] = elem + 1
		} else {
			ret[v] = 1
		}
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
