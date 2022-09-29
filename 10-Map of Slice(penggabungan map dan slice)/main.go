package main

import "fmt"

func main() {
	//membuat variabel slice berisi map
	nilaiSiswa := []map[string]string{
		{"nama": "Agung", "nilai": "A"},
		{"nama": "Joko", "nilai": "B"},
	}

	fmt.Println(nilaiSiswa)

	//for each
	for _, isi := range nilaiSiswa{
		fmt.Println(isi["nama"] , "--" , isi["nilai"])
	}
}