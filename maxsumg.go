package main

import (
	"fmt"
	"math"
)

func main() {
	//arr := []int{2, 1, 5, 2, 3, 2}
	//arr := []int{2, 1, 5, 2, 8}
	arr := []int{3, 4, 1, 1, 6}
	sm := 8

	window_p := 0
	window_s := 0
	min_len := math.MaxInt
	for window_e := 0; window_e < len(arr); window_e++ {
		window_p = window_p + arr[window_e]

		for window_p >= sm {
			min_len = min(min_len, window_e-window_s+1)
			window_p = window_p - arr[window_s]
			window_s++
		}
	}

	fmt.Println(min_len)
}
