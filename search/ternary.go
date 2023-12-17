package search

func TernarySearch(arr []int, target int) int {
	return ternarySearchHelper(arr, 0, len(arr)-1, target)
}

func ternarySearchHelper(arr []int, left, right, target int) int {
	if right >= left {
		mid1 := left + (right-left)/3
		mid2 := right - (right-left)/3

		if arr[mid1] == target {
			return mid1
		}

		if arr[mid2] == target {
			return mid2
		}

		if target < arr[mid1] {
			return ternarySearchHelper(arr, left, mid1-1, target)
		} else if target > arr[mid2] {
			return ternarySearchHelper(arr, mid2+1, right, target)
		} else {
			return ternarySearchHelper(arr, mid1+1, mid2-1, target)
		}
	}
	return -1
}
