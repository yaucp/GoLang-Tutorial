package main

import (
	"fmt"
	"math/rand"
)

func init() {
	//	This function runs first to set up state... in every source files
	//	init runs after all var. declaration in the package have run after initialization

	fmt.Println("this functions run first before main hahaha")
}

func main() {

	c := 42
	//	If statements
	if c > 43 {
		fmt.Println("Hey this is impossible!")
	} else {
		fmt.Println("Hello this is the proper way!")
	}

	//	statement; statement idiom
	//	scope of y is only limited here (in if-block)
	if y := 2 * rand.Intn(c); y > 10 {
		fmt.Printf("y's randomly generated num (%v max) is larger than 10\n", c)
	}

	num := rand.Intn(200)
	//	switch statements
	switch {
	case num > 150:
		fmt.Println("Larger than 150 damn!")
	//	No fallthrough! No need to add break! Nice!
	case num == 150:
		fmt.Println("hiiii")
	default:
		fmt.Println("rest of the cases are here....")
	}
	//	OR
	switch num {
	case 100:
		fmt.Println(100)
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("rest")
	}
	//	Add fallthrough key word to run the next case ....

	//	For loop
	//	Refer to effective go
	// Basic C-like for loop:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//	for = while (inf. loop just put for)
	for c < 45 {
		fmt.Println(c)
		c++
	}

	//	for range loops
	xi := []int{41, 42, 4, 1, 100}
	for idx, val := range xi {
		fmt.Printf("%v index %v num\n", idx, val)
	}
	//	map
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	//Map DONT return things in order!
	for key, val := range m {
		fmt.Printf("%v key %v val\n", key, val)
	}
}
