package main

import (
	"fmt"
	//	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	fmt.Printf("%+v\n", p)

	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 14) // binary digit

	fmt.Printf("%c\n", 97) // ascii code

	fmt.Printf("%x\n", 456) // hex

	fmt.Printf("%f\n", 78.9)

	fmt.Printf("%p\n", &p)
}
