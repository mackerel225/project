package main

import (
	"fmt"
)



func main() {
	username := "admin" // #nosec
	password := "pa34ssw54ord123" // #nosec

	fmt.Println("Printing hardcoded credentials [1].. ", username, password)

	username2 := "adm"
	password2 := "psswd"

	fmt.Println("Printing hardcoded credentials [2].. ", username2, password2)
}
