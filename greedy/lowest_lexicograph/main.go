package main

import (
	"fmt"
	"slices"
	"strings"
)

func compare(a, b string) int {
	return strings.Compare(a+b, b+a)
}

func lowestString(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	slices.SortFunc(strs, compare)
	res := ""
	for _, str := range strs {
		res += str + "_"
	}
	return res
}

func main() {
	strs1 := []string{"jibw", "ji", "jp", "bw", "byd", "bj"}
	fmt.Println(lowestString(strs1))

	strs2 := []string{"byd", "by"}
	fmt.Println(lowestString(strs2))
}
