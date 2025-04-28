package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter two values to calculate with operator (e.g., 5+2):")
	input, _ := reader.ReadString('\n')


	input = strings.TrimSpace(input)
	
	operator := strings.IndexAny(input, "+-*/")
	
	if operator == -1 {
		fmt.Println("Invalid operator. Please use +, -, *, or /.")
		return
	}
	// split the input string into two parts: left and right of the operator
	left := input[:operator]
	right := input[operator+1:]

	// get the operator character
	op := string(input[operator])
	// check if left and right are valid numbers
	if _, err := strconv.Atoi(left); err != nil {
		fmt.Println("Invalid left operand. Please enter a valid number.")
		return
	}
	if _, err := strconv.Atoi(right); err != nil {
		fmt.Println("Invalid right operand. Please enter a valid number.")
		return
	}
	// convert left and right to integers
	leftInt, _ := strconv.Atoi(left)
	rightInt, _ := strconv.Atoi(right)

	switch op {
	case "+":
		result := leftInt + rightInt
		fmt.Printf("Addition: %d + %d = %d\n", leftInt, rightInt, result)
	case "-":
		result := leftInt - rightInt
		fmt.Printf("Subtraction: %d - %d = %d\n", leftInt, rightInt, result)
	case "*":
		result := leftInt * rightInt
		fmt.Printf("Multiplication: %d * %d = %d\n", leftInt, rightInt, result)
	case "/":
		if rightInt == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
			return
		}
		result := leftInt / rightInt
		fmt.Printf("Division: %d / %d = %d\n", leftInt, rightInt, result)
	default:
		fmt.Println("Unknown operator:", op)
	}

}
