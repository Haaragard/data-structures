package main

import "fmt"

func main() {
	naturalNumbersFor(25)
	naturalNumbersCalc(25)
}

func naturalNumbersFor(number int) {
	var sum int
	for i := 1; i <= number; i++ {
		sum += i
		fmt.Printf("FOR - Natural number iteration %d, sum: %d\n", i, sum)
	}
}

func naturalNumbersCalc(number int) int {
	var sum int = number * (number + 1) / 2
	fmt.Printf("CALC - Natural number sum: %d\n", sum)
	return sum
}
