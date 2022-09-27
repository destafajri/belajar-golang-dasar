package main

import "fmt"

func main() {

	sekarang := "Golang Is The Best Language"

	//mencetak index ganjil
	for index, huruf := range sekarang {
		if index % 2 != 0 {
		fmt.Println(string(huruf))
		}
	}

	//mencetak index genap
	for index, huruf := range sekarang {
		if index % 2 == 0 {
		fmt.Println(string(huruf))
		}
	}

	//mencetak huruf vocal
	for _, huruf := range sekarang {
		if string(huruf) == "a" || string(huruf) == "i" || string(huruf) == "u" || string(huruf) == "e" || string(huruf) == "o" {
		fmt.Print(string(huruf))
		}
	}
}