package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "hello沙河小王子hah嘿"

	s2 := "abcdefghijklmnopqrstuvwxyz"

	num := 0
	for _, c := range s1 {
		if !strings.ContainsRune(s2, c) {
			num++
		}
	}

	fmt.Println(num)
}
