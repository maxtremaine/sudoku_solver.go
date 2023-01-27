package main

import "sort"

func lastIndex(arr []int, item int) int {
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

func uniqueNumbers(arr []int) []int {
	var output []int
	for i, value := range arr {
		if i == lastIndex(arr, value) {
			output = append(output, value)
		}
	}
	sort.Ints(output)
	return output
}