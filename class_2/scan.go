package main

import (
	"bufio"
	"fmt"
	"os"
)

func inputScanner() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a value:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Input value:", input)
}
