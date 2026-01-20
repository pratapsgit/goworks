package main

//Given an array of positive numbers and a positive number K, find the maximum sum of any contiguous subarray of size K.
import (
	"fmt"
	"math"
)

func main() {
	//arr := []int{2, 1, 5, 1, 3, 2}
	//sz := 3
	arr := []int{2, 3, 4, 1, 5}
	sz := 2

	window_p := 0
	max_sum := math.MinInt
	window_s := 0
	for window_e := 0; window_e < len(arr); window_e++ {
		window_p = window_p + arr[window_e]

		if window_e >= sz-1 {
			max_sum = max(max_sum, window_p)

			window_p = window_p - arr[window_s]
			window_s = window_s + 1
		}
	}

	fmt.Println(max_sum)
}
