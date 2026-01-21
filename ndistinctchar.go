package main

// Given a string, find the length of the longest substring in it with no more than K distinct characters.

// You can assume that K is less than or equal to the length of the given string.
import (
	"fmt"
	"math"
)

func main() {
	//str := "araaci"
	str := "cbbebi"
	k := 3

	r := []rune(str)

	j := 0
	max_len := math.MinInt

	h := make(map[rune]int)
	for i := 0; i < len(r); i++ {
		if _, ok := h[r[i]]; !ok {
			h[r[i]] = 0
		}
		h[r[i]]++

		for len(h) > k {
			val := r[j]
			h[val]--
			if h[val] == 0 {
				delete(h, val)
			}

			j++
		}
		max_len = max(max_len, i-j+1)
	}

	fmt.Println(max_len)
}
