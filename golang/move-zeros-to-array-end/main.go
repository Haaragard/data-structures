package main

import "fmt"

func main() {
	fmt.Println(moveZeros([]int{0, 0, 0, 0, 1, 2, 0, 3, 0, 4, 0, 0, 0, 5, 5, 1, 2, 7, 33, 0, 3, 0, 4, 5, 0, 0, 0, 3, 3, 3, 3, 1, 2, 3, 0}))
}

func moveZeros(arr []int) []int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 && arr[j] == 0 {
			arr[i], arr[j] = arr[j], arr[i]
		}

		if arr[j] != 0 {
			j++
		}
	}

	return arr
}
