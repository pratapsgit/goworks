package main

// You are visiting a farm that has a single row of fruit trees arranged from left to right. The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.

// You want to collect as much fruit as possible. However, the owner has some strict rules that you must follow:

// You only have two baskets, and each basket can only hold a single type of fruit. There is no limit on the amount of fruit each basket can hold.
// Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right. The picked fruits must fit in one of your baskets.
// Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
// Given the integer array fruits, return the maximum number of fruits you can pick.

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 3, 2, 2}

	window_s := 0

	max_fruits := math.MinInt
	mhash := make(map[int]int)
	for window_e := 0; window_e < len(arr); window_e++ {
		if _, ok := mhash[arr[window_e]]; !ok {
			mhash[arr[window_e]] = 0
		}
		mhash[arr[window_e]]++

		for len(mhash) > 2 {
			e := arr[window_s]
			mhash[e]--

			if mhash[e] == 0 {
				delete(mhash, e)
			}

			window_s++
		}

		max_fruits = max(max_fruits, window_e-window_s+1)
	}

	fmt.Println(max_fruits)
}
