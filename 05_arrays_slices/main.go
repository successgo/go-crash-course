package main

import "fmt"

func main() {
	var linux [2]string

	linux[0] = "Debian"
	linux[1] = "CentOS"

	fmt.Println(linux)
	fmt.Println(linux[1])

	mobile := []string{"Android", "iOS", "Ubuntu"}

	fmt.Println(mobile)
	fmt.Println(len(mobile))
	fmt.Println(mobile[1:3])
}
