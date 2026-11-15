package main

import "fmt"

func main() {
	var reverseArray = reverseArray([]int{1, 2, 3, 4, 5, 6})
	printArray(reverseArray)
}

func reverseArray(arr []int) []int {
	start := 0
	end := len(arr) - 1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

	return arr
}

func printArray(list []int) {
	for _, item := range list {
		fmt.Printf("%d\t", item)
	}
	fmt.Print("\n")
}
