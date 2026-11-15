package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 55, 123123}
	printArray(slice)

	// var array = [2]int{1, 2}
	// printArray(array) : Error

}

func printArray(list []int) {
	for _, item := range list {
		fmt.Printf("%d\t", item)
	}
	fmt.Print("\n")
}
