package main

import "fmt"

const (
	_ = iota // c0 == 0
	a
	b
	c
	d
	e
	f
)

type ByteSize int //creating our own type

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // like initialise the starting format
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)

	fmt.Printf("%v lmao wtf %T\n", a, a)
}
