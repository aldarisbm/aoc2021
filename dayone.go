package main

import (
	"strconv"
	"strings"
)

func getDayOneResult() int {
	input, err := getDayInput(1)
	if err != nil {
		panic(err)
	}
	inputValues, err := getDayOneValues(input)
	if err != nil {
		panic(err)
	}
	var count = 0
	for i := 0; i < len(inputValues)-1; i++ {
		if inputValues[i] < inputValues[i+1] {
			count++
		}
	}
	return count
}

func getDayOneValues(input string) ([]int, error) {
	inputSplit := strings.Split(input, "\n")
	var intArray []int
	for _, element := range inputSplit {
		if len(element) > 0 {
			castValue, err := strconv.Atoi(element)
			if err != nil {
				return nil, err
			}
			intArray = append(intArray, castValue)
		}
	}
	return intArray, nil
}
