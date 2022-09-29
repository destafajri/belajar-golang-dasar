package main

import "fmt"

func main() {
	//array string
	var kata[5]string
	kata[0] = "Go"
	kata[4] = "Java"
	fmt.Println(kata)

	pemrograman := [5] string{"Python", "JavaScript"}
	fmt.Println(pemrograman)

	//array integer
	var angka[5] int
	angka[2] = 3
	angka[4] = 4
	fmt.Println(angka)

	//array dinamis
	dinamis := [...] int {
		4,
		2,
		3,
		5,
	}

	fmt.Println(dinamis)

	//mengetahui panjang array
	fmt.Println(len(kata))

	panjang := len(angka)
	fmt.Println(panjang)
	
	fmt.Println(len(dinamis))


}