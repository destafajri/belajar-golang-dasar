package main

import "fmt"

func main() {
	testcount(5)
}

//looping dengan recursive function
func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}