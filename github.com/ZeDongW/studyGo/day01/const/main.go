package main

import "fmt"

const pi = 3.14

const (
	statusOK = 200
	notFund  = 400
)

const (
	n1 = 100
	n2
	n3
)

const (
	a1 = iota
	a2
	a3
)

const a4 = iota

func main() {
	fmt.Println(pi)

    fmt.Println(statusOK)
	fmt.Println(notFund)

	fmt.Println("n1, ", n1)
	fmt.Println("n2, ", n2)
	fmt.Println("n3, ", n3)

	fmt.Println("a1, ", a1)
	fmt.Println("a2, ", a2)
	fmt.Println("a3, ", a3)

	fmt.Println("a4, ", a4)
}