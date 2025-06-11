package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <num1> <operator> <num2>")
		fmt.Println("Example: go run main.go 10 + 5")
		return
	}

	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	op := os.Args[2]
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: please enter valid numbers..")
		return
	}
	var result float64

	switch op {
	case "+":
		result = num1 + num2

	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: division by zero")
			return
		}
		result = num1 / num2
	case "%":
		result = float64(int(num1) % int(num2))
	default:
		fmt.Println("Error: Unsupported operators..")
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}
