package main

import "fmt"

func main() {
	// Declare h to hold a value of type Int <- good for declare and not assign concrete value yet (but with default val)
	// Int -> 0, String -> "", float -> 0.0, boolean -> false, otherwise -> nil
	var h int = 42
	fmt.Println(h)

	// Short declaration (auto infer type)
	// Only in function
	a, b, c, d := 32, 12, 98, 100
	fmt.Println(a, b, c, d)
}
