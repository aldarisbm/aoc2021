package main

import (
	"strconv"
	"strings"
)

func getDayOneResult() (int, int, error) {
	input, err := getDayInput(1)
	if err != nil {
		return 0, 0, nil
	}
	inputValues, err := getDayOneValues(input)
	if err != nil {
		return 0, 0, nil
	}
	var count = 0
	for i := 0; i < len(inputValues)-1; i++ {
		if inputValues[i] < inputValues[i+1] {
			count++
		}
	}
	var secondCount = 0
	for i := 0; i < len(inputValues)-4; i++ {
		firstGroup := inputValues[i] + inputValues[i+1] + inputValues[i+2]
		secondGroup := inputValues[i+2] + inputValues[i+3] + inputValues[i+4]
		if secondGroup > firstGroup {
			secondCount++
		}
	}
	return count, secondCount, nil
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
