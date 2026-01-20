package main

//Given an array, find the average of all contiguous subarrays of size K in it.
import (
	"fmt"
)

func calculateAverage(arr []int, k int) []float32 {

	window_p := 0
	window_s := 0
	result := make([]float32, len(arr)-k+1)

	for window_e := 0; window_e < len(arr); window_e++ {
		window_p = window_p + arr[window_e]

		if window_e >= k-1 {
			result[window_s] = float32(window_p) / float32(k)
			window_p = window_p - arr[window_s]
			window_s++
		}
	}

	return result
}

func main() {
	arr := []int{3, 1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 5

	result := calculateAverage(arr, k)

	for _, e := range result {
		fmt.Printf("%.2f, ", e)
	}
	fmt.Println()
}
