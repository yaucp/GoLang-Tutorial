package main

import "fmt"

func main() {
	//Array
	//NUMBERED sequence with VALUES of the SAME type only (size unchanged)
	//Mostly used for Go internals, generally not used

	//Array Literal
	a := [3]int{1, 2, 3}
	fmt.Println(a)

	b := [...]string{"Hello", "There"}
	fmt.Printf("Length: %v \t %v\n", len(b), b)

	var c [2]int
	c[0] = 10
	c[1] = 100
	fmt.Println(c)

	//Slice
	//Built on arrays but can expand/change sizes
	//Have length & capacity

	//Map
	//UNORDERED, key-element

}
