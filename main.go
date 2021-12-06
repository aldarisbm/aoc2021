package main

import (
	"fmt"
)

func main() {
	firstDayOneResult, secondDayOneResult, err := getDayOneResult()
	if err != nil {
		panic(err)
	}
	fmt.Printf("First Day One Result: %d, Second Day One Result: %d", firstDayOneResult, secondDayOneResult)

	// dayTwoResult, err := getDayTwoResult()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Day Two Result: %d", dayTwoResult)
}
