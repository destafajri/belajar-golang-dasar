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

	information := userInfo(user1)
	fmt.Println(information)

}

//fuction menggunakan parameter struct
func userInfo(info user) string{
	result := fmt.Sprintf("Name : %s %s, Address : %s ", info.firstName, info.lastName, info.address)
	return result
}