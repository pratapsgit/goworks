package main

//Given an array of positive numbers and a positive number S, find the pairs with sum S with no duplicates.
import (
	"fmt"
	"slices"
)

func aPairWithSum(arr []int, start int, end int, target int) ([][]int, error) {
	s := start
	e := end

	var result [][]int

	slices.Sort(arr)

	for s < e {
		tsum := arr[s] + arr[e]

		if tsum == target {
			result = append(result, []int{arr[s], arr[e]})
			s++
			e--

			for s < end && arr[s] == arr[s-1] {
				s++
			}

			for e > 0 && arr[e] == arr[e+1] {
				e--
			}
		} else if tsum > target {
			e--
		} else {
			s++
		}
	}

	return result, nil
}

func main() {
	arr := []int{2, 5, 1, 4, 6, 8, 7, 9, 2, 5, 1, 4, 6, 3, 1}
	target := 8

	fmt.Println(arr, " to find target sum ", target)
	result, err := aPairWithSum(arr, 0, len(arr)-1, target)

	if err == nil {
		for _, r := range result {
			fmt.Printf("{%d, %d}\n", r[0], r[1])
		}
	}
}
