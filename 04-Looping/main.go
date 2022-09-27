package main

import "fmt"

func main() {
	//looping for
	for i:= 1; i <= 100; i++ {
		fmt.Println("Lagi belajar Go nihh")
	}

	//for each
	lagiapa := "Saya lagi belajar Golang"

	for index, letter:= range lagiapa{
		fmt.Println("index",index, "letter", string(letter))
	}

	//skip index dari perulangan for each
	for _, letter:= range lagiapa{
		fmt.Println("letter", string(letter))
	}
}