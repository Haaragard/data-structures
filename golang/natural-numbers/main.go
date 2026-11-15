package main

import (
	"fmt"
	"time"
)

func main() {
	const value int = 999999999
	naturalNumbersIterate(value)
	naturalNumbersCalc(value)
}

func naturalNumbersIterate(number int) {
	var sum int

	timeStart := time.Now()
	for i := 1; i <= number; i++ {
		sum += i
	}
	timeEnd := time.Now()
	fmt.Printf("ITER -----------\n")
	fmt.Printf("ITER - Result: %d\n", sum)
	fmt.Printf("ITER - Duration in milliseconds: %d\n", timeEnd.Sub(timeStart).Milliseconds())
	fmt.Printf("ITER -----------\n")
}

func naturalNumbersCalc(number int) {
	timeStart := time.Now()
	var sum int = number * (number + 1) / 2
	timeEnd := time.Now()
	fmt.Printf("CALC -----------\n")
	fmt.Printf("CALC - Result: %d\n", sum)
	fmt.Printf("CALC - Duration in milliseconds: %d\n", timeEnd.Sub(timeStart).Milliseconds())
	fmt.Printf("CALC -----------\n")
}
