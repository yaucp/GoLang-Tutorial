package main

import "fmt"

func main() {
	// No casting only conversion in GoLang!

	y := 42   // int
	z := 42.0 // float64
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	/*
		// this does not work!
		// in go you can't take a VALUE that is float32 and store it
		// in a variable that is declared to hold a VALUE of float64
		z = m
		fmt.Printf("%v of type %T \n", z, z)
	*/

	// this does work
	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)

	// GoLang requires EXPLICIT conversion during arithmetic operations

	// struct (object in JS)
	// constructing struct: composition
}
