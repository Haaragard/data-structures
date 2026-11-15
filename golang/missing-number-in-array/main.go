package main

import "fmt"

func main() {
	var array = []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0}

	missingNumber := findMissingNumber(array)
	fmt.Println(missingNumber)
}

func findMissingNumber(arr []int) int {
	arrLen := len(arr)
	arrTotalSum := arrLen * (arrLen + 1) / 2

	for _, value := range arr {
		arrTotalSum -= value
	}

	return arrTotalSum
}
