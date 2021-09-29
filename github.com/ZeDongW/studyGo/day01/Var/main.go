package main

import "fmt"

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "ZeDongW"
	age = 25
	isOk = false
	fmt.Printf("name: %s, age: %d, isOk: %t", name, age, isOk)
	fmt.Println()

	var s1 string = "wwx"
	fmt.Println(s1)

	var s2 = 20 
	fmt.Println(s2)

	s3 := false
	fmt.Println(s3)
}
