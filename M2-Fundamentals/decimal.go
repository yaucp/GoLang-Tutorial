package main

import (
	"fmt"
)

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	// %v for just the value in a default format
	fmt.Printf("%b %b %b %b %b %b\n", a, b, c, d, e, f)
	fmt.Printf("%x %x %x %x %x %x \n", a, b, c, d, e, f)
}
