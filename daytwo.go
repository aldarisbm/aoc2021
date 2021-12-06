package main

import (
	"strconv"
	"strings"
)

type DirectionStep struct {
	Direction string
	Steps     int
}

func getDayTwoResult() (int, error) {
	input, err := getDayInput(2)
	if err != nil {
		return 0, err
	}
	inputValues, err := getDayTwoValues(input)
	if err != nil {
		return 0, err
	}
	horizontalPos, depthPos := getHorizontalAndDepth(inputValues)
	return horizontalPos * depthPos, nil
}

func getDayTwoValues(input string) ([]DirectionStep, error) {
	inputSplit := strings.Split(input, "\n")
	var stringArray []string
	for _, element := range inputSplit {
		if len(element) > 0 {
			stringArray = append(stringArray, element)
		}
	}
	directionStepArray := []DirectionStep{}
	for _, element := range stringArray {
		arr := strings.Split(element, " ")
		castValue, err := strconv.Atoi(arr[1])
		if err != nil {
			return nil, err
		}
		dirStepObj := DirectionStep{Direction: arr[0], Steps: castValue}
		directionStepArray = append(directionStepArray, dirStepObj)
	}
	return directionStepArray, nil
}

func getHorizontalAndDepth(input []DirectionStep) (horizontal int, depth int) {
	for _, elem := range input {
		switch elem.Direction {
		case "forward":
			horizontal = horizontal + elem.Steps
		case "down":
			depth = depth + elem.Steps
		case "up":
			depth = depth - elem.Steps
		}
	}
	return
}
