package main

import (
	"fmt"
	"errors"
)

func main() {
	//penjumlahan array(soal-1)
	scores :=[]int {10, 5, 8, 9, 7}
	hasil := sum(scores)
	fmt.Println("hasil penjumlahan array", hasil)

	//operasi matematika(soal-2)
	//tanpa error
	fmt.Println(calculation(5,5,"+"))

	//dengan error
	hasil, errorMsg := calculation(5,6,"=")
	fmt.Println("hasil cetak biasa",hasil, errorMsg)
	//mencetak dengan if condition
	if errorMsg != nil{
		fmt.Println(errorMsg.Error())
	}else{
		fmt.Println(hasil)
	}

	//tanpa error
	//dengan error
	hasil2, errorMsg2 := calculation(12,6,"*")
	//mencetak dengan if condition
	if errorMsg2 != nil{
		fmt.Println(errorMsg2.Error())
	}else{
		fmt.Println(hasil2)
	}
}

//fungsi penjumlahan array
func sum(nilai[]int)int{
	var jumlah int
	for i := 0; i < len(nilai); i++ {
		jumlah += nilai[i]
	}
	return jumlah
}

//fungsi aritmatika multiple return dengan error message
func calculation(nilai1, nilai2 int, operator string) (int, error){
	if operator == "+"{
		return nilai1 + nilai2, nil
	}else if operator == "-"{
		return nilai1 - nilai2, nil
	}else if operator == "*"{
		return nilai1*nilai2, nil
	}else if operator == "/" {
		return nilai1/nilai2, nil
	}else{
		return 0, errors.New("operator salah")
	}
}
