package calculator

import (
	"errors"
)

func Calculate(command string, inputX int, inputY int) (int, error) {
	var output int
	var err error
	switch command {
	case "add":
		output = inputX + inputY
	case "subtract":
		output = inputX - inputY
	case "multiply":
		output = inputX * inputY
	case "divide":
		output = inputX / inputY
	default:
		err = errors.New("Invalid command")
	}
	return output, err
}
