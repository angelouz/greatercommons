package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

//go has conversion, not casting
func main() {
	a = 42
	b = 43

	fmt.Println(a)
	fmt.Printf("%T\n\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n\n", b)

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n\n", a)
}
