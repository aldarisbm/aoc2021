package main

import (
	"fmt"
)

func main() {
	// firstDayOneResult, secondDayOneResult, err := getDayOneResult()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("First Day One Result: %d, Second Day One Result: %d", firstDayOneResult, secondDayOneResult)

	// firstDayTwoResult, secondDayTwoResult, err := getDayTwoResult()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("First Day Two Result: %d, Second Day Two Result: %d", firstDayTwoResult, secondDayTwoResult)

	firstDayThreeResult, secondDayThreeResult, err := getDayThreeResult()
	if err != nil {
		panic(err)
	}
	fmt.Printf("First Day Three Result: %d, Second Day Three Result: %d", firstDayThreeResult, secondDayThreeResult)
}
