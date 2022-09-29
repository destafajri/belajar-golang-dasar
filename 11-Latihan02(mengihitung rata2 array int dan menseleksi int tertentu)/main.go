package main

import "fmt"

func main() {
	nilai := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	//menghitumg rata2
	var tampung int
	for i := 0; i < len(nilai); i++ {
		tampung += nilai[i]
	}

	rata2 := float64(tampung)/float64(len(nilai))
	fmt.Println(rata2)

	//spesifik nilai
	var goodScore[] int
	for i := 0; i < len(nilai); i++ {
		if nilai[i] >= 90{
			goodScore = append(goodScore, nilai[i])
		}
	}

	fmt.Println(goodScore)

	for _, cetak := range goodScore{
		fmt.Println(cetak)
	}
}