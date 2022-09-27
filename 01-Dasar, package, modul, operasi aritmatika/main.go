package main

import (
	"belajar-golang-dasar/calculation"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	add := calculation.Add(8,9)
	multiple := calculation.Multiple(5,55)
	min := calculation.Minus(7,1)
	divide := calculation.Divide(8,2)

	fmt.Println("add ", add)
	fmt.Println("multiple ", multiple)
	fmt.Println("minus ", min)
	fmt.Println("divided ", divide)
}