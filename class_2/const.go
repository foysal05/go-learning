package main

import "fmt"

func constExample() {

	const a, b = 5, 7
	const cp = a + b // Syntax for constant declaration

	cal := cp + 5
	fmt.Printf("Const value :%d\n", cal) // Address of a

	const (
		initial = 1 << iota // 1 << 0 = 1
		second              // 1 << 1 = 2
		third               // 1 << 2 = 4
		fourth              // 1 << 3 = 8

	)
	const (
		_        = 4 << iota // 4 << 0 = 4
		second_1             // 4 << 1 = 8
		third_1              // 4 << 2 = 16
		fourth_1             // 4 << 3 = 32

	)

	fmt.Println("Const value : ", initial, second, third, fourth) // Address of a
	fmt.Println("Const value : ", second_1, third_1, fourth_1)    // Address of a
	inputScanner()

}
