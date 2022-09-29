package main

import "fmt"

func main() {
	luas, keliling :=luasdankeliling(5,6)
	fmt.Println(luas, keliling)
}

//function multiple return dengan predefined return
func luasdankeliling(p,l int) ( luas int, keliling int){
	if p !=0 && l != 0{
		luas = p*l
		keliling = 2*(p+l)
		return
	}else {
		return 0, 0
	}
}