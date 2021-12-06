package main

import (
	"strconv"
	"strings"
)

func getDayThreeResult() (int, int, error) {
	input, err := getDayInput(3)
	if err != nil {
		return 0, 0, err
	}
	inputValues, err := getDayThreeValues(input)
	if err != nil {
		return 0, 0, err
	}
	firstResult := getPowerConsumption(inputValues)
	return firstResult, 0, nil
}

func getDayThreeValues(input string) ([]int, error) {
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

func getPowerConsumption(input []int) int {

	return 0
}
