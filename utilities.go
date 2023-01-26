package main

func lastIndex(arr []int, item any) int {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] == item {
			return i
		}
	}

	return 0
}

func hasDuplicates(arr []int) bool {
	for i, item := range arr {
		if lastIndex(arr, item) != i {
			return true
		}
	}

	return false
}