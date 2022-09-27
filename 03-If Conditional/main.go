package main

import "fmt"

func main() {

	//if Statement
	age := 20
	status := "jomblo"

	if age >= 20 {
		fmt.Println("Sudah Dewasa")
	}

	//if else Statement
	if  age >= 21 {
		fmt.Println("Boleh minum alkohol")
	} else{
		fmt.Println("Belum boleh maboook")
	}

	//nested if

	if age >17 {
		if status == "jomblo" {
			fmt.Println("Anda menyedihkan")
		}
	}
}