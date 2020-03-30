package main

import (
	"fmt"
	"math"

	"github.com/successgo/go-crash-course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(3.7))
	fmt.Println(math.Ceil(3.7))
	fmt.Println(math.Pow(2, 10))
	fmt.Println(math.Sqrt(9))
	fmt.Println(strutil.Reverse("hello"))
}