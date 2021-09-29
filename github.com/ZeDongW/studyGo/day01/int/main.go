package main

import "fmt"

func main() {
	i1 := 101
	fmt.Printf("%b\n", i1)
	fmt.Printf("%o\n", i1)
	fmt.Printf("%x\n", i1)

	i2 := 077
	fmt.Printf("%d\n", i2)

	i3 := 0xfffff
	fmt.Printf("%o\n", i3)

	i4 := int8(4)
	fmt.Printf("%T\n", i4)

	i5 := int64(32323)
	fmt.Printf("%T\n", i5)
}