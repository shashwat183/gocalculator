package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/shashwat183/gocalculator/calculator"
	"log"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Please provide a command and two number like so - add <num1> <num2>")
	}
	command := os.Args[1]
	inputX, inputXErr := strconv.Atoi(os.Args[2])
	inputY, inputYErr := strconv.Atoi(os.Args[3])
	if inputXErr != nil || inputYErr != nil {
		log.Fatal("Invalid numeric arguments, please check ur input")
	}
	output, err := calculator.Calculate(command, inputX, inputY)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result = %v\n", output)
}
