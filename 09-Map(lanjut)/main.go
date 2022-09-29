package main

import "fmt"

func main() {
	//for each pada map
	myMap := map[int]string{}
	myMap[0] = "rubby"
	myMap[2] = "Golang"
	myMap[7] = "Java"

	fmt.Println(myMap)

	for key, nilai := range myMap{
		fmt.Println(key, nilai)
	}

	fmt.Println("=======")
	
	//menghapus isi map
	//delete(namaVariabel, key)
	delete(myMap, 2)
	
	for key, nilai := range myMap{
		fmt.Println(key, nilai)
	}

	//menambahkan kembali
	myMap[2] = "python"

	fmt.Println("=======")
	//menampilkan nilai atau isi map
	tampil := myMap[2]
	tampil2 := myMap[7]

	fmt.Println(tampil)
	fmt.Println(tampil2)

	fmt.Println("=======")
	//mengecek secara boolean ketersediaan nilai
	nilai, isAvail := myMap[0]
	_, isAvailable := myMap[1]
	
	fmt.Println(nilai)
	fmt.Println(isAvail)
	fmt.Println(isAvailable)
}