package main

import "fmt"

//membuat struct
type Mahasiswa struct{
	ID int
	firstName string
	lastName string
	isActive bool
}

func main() {
	//membuat object
	mahasiswa := Mahasiswa{}

	//mengisi object
	mahasiswa.ID = 1
	mahasiswa.firstName = "Otong"
	mahasiswa.lastName = "Surotong"
	mahasiswa.isActive = true

	fmt.Println(mahasiswa)
	fmt.Println(mahasiswa.firstName)

	//mengisi object secara langsung
	mahasiswa2 := Mahasiswa{ID:2, firstName: "bakso", lastName: "goreng", isActive: false}

	fmt.Println(mahasiswa2)
	fmt.Println(mahasiswa2.isActive)
}