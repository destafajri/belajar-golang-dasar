package main

import "fmt"

func main() {

	//dinamis array
	var anggota []string

	//menambahkan nilai
	anggota = append(anggota, "anggota 1")

	fmt.Println(anggota)

	//dinamis array langung mengisi nilai
	manusia :=[]string {"gogon", "andrew"}

	fmt.Println(manusia)

	//print for each
	for  _, nama := range manusia{
		fmt.Println(nama)
	}
}