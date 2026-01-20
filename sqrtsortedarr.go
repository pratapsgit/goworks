package main

import (
	"fmt"
	"slices"
)

func sortedSqrt(arr []int) []int {
	slices.Sort(arr)

	for i, e := range arr {
		fmt.Println(i, " ", e)
	}

	return []int{}
}

func main() {
	arr := []int{-3, 1, 5, 2, -2, 6, 4, 7, -1}

	result := sortedSqrt(arr)

	fmt.Println(result)
}
