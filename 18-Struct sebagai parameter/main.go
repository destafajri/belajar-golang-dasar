package main

import "fmt"

type user struct{
	ID int
	firstName string
	lastName string
	address string
	status bool
}

func main() {
	user1 := user{1, "nana", "nini", "jakarta",true}

	info := userInfo(user1)
	fmt.Println(info)

}

//fuction menggunakan parameter struct
func userInfo(info user) string{
	result := fmt.Sprintf("Name : %s %s, Adress : %s ", info.firstName, info.lastName, info.address)
	return result
}