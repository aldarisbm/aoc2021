package main

import (
	"fmt"
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
	firstResult, err := getPowerConsumption(inputValues)
	if err != nil {
		return 0, 0, err
	}
	return firstResult, 0, nil
}

func getDayThreeValues(input string) ([]string, error) {
	inputSplit := strings.Split(input, "\n")
	var stringArray []string
	for _, element := range inputSplit {
		if len(element) > 0 {
			stringArray = append(stringArray, element)
		}
	}
	return stringArray, nil
}

func getPowerConsumption(input []string) (int, error) {
	var intArray []int
	for i := 0; i < len(input[0]); i++ {
		intArray = append(intArray, 0)
	}
	for _, element := range input {
		for index, letter := range element {
			switch letter {
			case '0':
				intArray[index]--
			case '1':
				intArray[index]++
			}
		}
	}
	gammaString := ""
	epsilonString := ""
	for i := 0; i < len(intArray); i++ {
		switch {
		case intArray[i] > 0:
			gammaString += "1"
			epsilonString += "0"
		case intArray[i] < 0:
			gammaString += "0"
			epsilonString += "1"
		}
	}
	fmt.Printf("Gamma string: %s\n", gammaString)
	fmt.Printf("Epsilon string: %s\n", epsilonString)
	gammaDecimal, err := strconv.ParseInt(gammaString, 2, 64)
	if err != nil {
		return 0, err
	}
	epsilonDecimal, err := strconv.ParseInt(epsilonString, 2, 64)
	if err != nil {
		return 0, err
	}
	return int(gammaDecimal) * int(epsilonDecimal), nil
}
