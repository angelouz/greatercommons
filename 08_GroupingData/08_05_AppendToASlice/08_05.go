package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	// "..." means variadic ->
	y := []int{234, 456, 648, 987}
	x = append(x, y...)
	fmt.Println(x)
}
