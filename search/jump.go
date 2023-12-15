package search

import (
	"math"
)

func JumpSearch(arr []int, target int) int {
	n := len(arr)
	blocksize := int(math.Sqrt(float64(n)))
	step := blocksize
	prev := 0

	for (step < n) && (arr[step-1] < target) {
		prev = step
		step += blocksize
		if prev >= n {
			return -1
		}
	}

	indexInBlock := LinearSearch(arr[prev:min(step, n)], target)

	if indexInBlock != -1 {
		return prev + indexInBlock
	}
	return -1
}
