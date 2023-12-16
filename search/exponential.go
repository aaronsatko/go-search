package search

func ExponentialSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	if arr[0] == target {
		return 0
	}

	index := 1
	for index < n && arr[index] <= target {
		index *= 2
	}

	low := index / 2
	high := min(index, n)
	result := BinarySearch(arr[low:high], target)

	if result != -1 {
		return result + low
	}
	return -1
}
