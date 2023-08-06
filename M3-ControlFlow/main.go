package main

import (
	"fmt"
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
}
