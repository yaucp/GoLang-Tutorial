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

	var c [2]int //Length of array is integrated into type as well
	c[0] = 10
	c[1] = 100
	fmt.Printf("%T \t%v\n", c, c)

	//Slice
	//Built on arrays but can expand/change sizes
	//Pointers (referencing) to the underlying array
	//Make changes to the underlying array automatically
	//Have length & capacity

	xs := []string{"hello", "golang"} //No number in the bracket (length)
	fmt.Println(xs, len(xs))

	for idx, val := range xs {
		fmt.Printf("%v - %v\n", idx, val)
	}

	//Append to slice
	xs = append(xs, "wassupppp", "woahhhhh")
	fmt.Println(xs)

	//Slicing the slice...
	//Same with python = [inclusive:exclusive]

	//Deleting from a slice
	//Removing "wassup"
	//... (similar to js) expand! variadic
	xs = append(xs[:2], xs[3:]...)
	fmt.Println(xs)

	//Make slices, allocates!
	//make(type, length, capacity)
	//Length: initially how many elements it has currently
	//Capacity: how many value to hold (auto expand when length exceeds the old capacity)
	ys := make([]int, 0, 10)
	fmt.Println(ys, len(ys), cap(ys))

	//2D Slice
	ppl1 := []string{"James", "Student"}
	ppl2 := []string{"Thomas", "Teacher"}
	yys := [][]string{ppl1, ppl2}
	fmt.Println(yys)

	//Copy (clone) array
	pplCopy := make([]string, 2)
	//copy(dst, src)
	copy(pplCopy, ppl1)
	pplCopy[0] = "Rachel"
	fmt.Println(ppl1, pplCopy)

	//GoLang function pass by value
	//func randomFunc(people []string)
	//Calling randomFunc(ppl1) ==> people := pp1 (same underlying array)
	//Can fix it by creating a new slice and copy the old one into the new variable

	//Map
	//UNORDERED, key-element

}
