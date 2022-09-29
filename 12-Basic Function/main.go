package main

import "fmt"

/* dalam fungsi terdapat
1.parameter/input
2.process
3.output/return/nilai
*/

func main() {
	//memanggil fuction tanpa parameter
	printMyFunction()
	
	//memanggil fuction dengan parameter dengan process tanpa nilai
	printMyFunction2("function dengan parameter tanpa return/output")

	//memanggil fuction dengan parameter dengan process dan nilai/output/return
	contoh := printMyFunction3("function dengan parameter,proses, dan return/output")
	fmt.Println(contoh)
	contohKosong := printMyFunction3("")
	fmt.Println(contohKosong)

	//math fuction
	hasil := add(1,2)
	fmt.Println(hasil)
}


//contoh function sederhana tanpa input/parameter dan output/pengembalian nilai
func printMyFunction() {
	fmt.Println("Function sederhana tanpa parameter")
}

//contoh function dengan parameter dan proses
func printMyFunction2(kata string){
	fmt.Println(kata)
}

//contoh function dengan input, proses, dan output/return
func printMyFunction3(kata string) string{
	sentence := kata
	if sentence == "" {
		sentence = "alert! tidak boleh kosong/harus diisi"
		return sentence
	} else {
		return kata
	}
}

//contoh function return value untuk operasi matematika
func add(nilai1, nilai2 int) int{
	tambah := nilai1+nilai2
	return tambah
}
