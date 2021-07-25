package main

import "fmt"

func main() {
	a := make(map[string]int)

	a["k1"] = 7
	a["k2"] = 13

	fmt.Println("map:", a)

	v1 := a["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(a))

	_, prs := a["k2"]
	fmt.Println("prs:", prs)

	_, prs = a["invalid"]
	fmt.Println("prs2:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
