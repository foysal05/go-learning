package main

import "fmt"

func main() {
	//Pointer example

	a := 10
	// b := &a // b is a pointer to a
	b := &a // Short syntax for pointer declaration
	// var b *int =&a   // Full syntax for pointer declaration
	a = 20

	fmt.Println("Value of a:", *b) // * is called dereferencing, it gets the value of a pointer

	constExample()
}
