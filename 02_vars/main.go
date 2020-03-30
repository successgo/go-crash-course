package main

import "fmt"

func main()  {
	// var name = "Success"
	var age int32 = 27
	var isCool = true
	isCool = false

	name, email := "Success", "successgao@protonmail.ch"
	size := 1.7

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", size)
}
