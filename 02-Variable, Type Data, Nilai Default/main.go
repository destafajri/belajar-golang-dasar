package main

import "fmt"

func main() {
	//type data string
	var name string = "Desta"
	fmt.Println("Stirng : ", name)

	//int
	var angka int = 100
	fmt.Println("Integer : ",angka)

	//float
	var pecahan float32 = 3.4444
	fmt.Println("Float32 : ",pecahan)

	//boolean
	var bol bool = true
	fmt.Println("boolean :",bol)

	//contoh nilai default
	var defaults string
	fmt.Println(defaults)

	//mengisi nilai variable
	defaults = "mengisi"
	fmt.Println(defaults)

	//tanpa inisiasi type data
	age := 20
	fmt.Println("integer tanpa inisiasi type data :",age)

	//reasign
	age = 30
	fmt.Print("mengisi ulang variable age ", age)
	
}