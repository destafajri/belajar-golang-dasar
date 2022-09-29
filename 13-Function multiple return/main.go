package main
import "fmt"

func main() {
	//memasukan parameter
	luas1, keliling := luasdankeliling(2,3)
	fmt.Println(luas1, keliling)

	luas2, _ := luasdankeliling(5,3)
	fmt.Println(luas2)

}

//function multiple return
func luasdankeliling(p,l int) (int, int){
	if p !=0 && l != 0{
		luas := p*l
		keliling := 2*(p+l)
		return luas, keliling
	}else {
		return 0, 0
	}
}
