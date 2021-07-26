package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	f := rect{width: 10, height: 5}
	fmt.Println("area:", f.area())
	fmt.Println("perim:", f.perim())

	rp := &f
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
