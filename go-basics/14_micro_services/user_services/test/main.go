package main

import (
	"14_micro_services/user_services/util"
	"fmt"
)

func main() {
	finalPassword := util.GeneratePassword("12345678 ")

	// $pbkdf2-sha512$49eZmV7IIkdb3a2104b9c1fba8102b402dba9482870fea84d0929b3b25274ef5f2d0edc65fd4ce0810bf89dd7c551d2a52356c6d29548f
	// $pbkdf2-sha512$VE4FEbVdmXe8616415bdf0eb0355b8242f63e7ac304dc109bb807d28a2c63a19058762e3482168d1d29c1a0b27096f97e20b5a2c1a5266
	fmt.Println(finalPassword)

	isPass := util.VerifyPassword("12345678 ", finalPassword+"1")

	fmt.Println(isPass)
}
