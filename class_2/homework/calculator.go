package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter two values to calculate (e.g., 5+2):")
	input, _ := reader.ReadString('\n')
	

	 fmt.Println("Input value:", input)
}
