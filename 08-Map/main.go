package main

import "fmt"

func main() {
	//inisiasi variabel
	var myMap map[int]string

	//inisiasi map
	myMap = map[int]string{}

	//megisi nilai map
	//map[key] = value
	myMap[0] = "Bruno"
	myMap[20]= "Rafaela"

	fmt.Println(myMap)

	//mendapatkan nilai suatu map meggunakan key ==> map[key]
	fmt.Println(myMap[20])
}