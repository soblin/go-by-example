package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	c := make([]string, len(s))
	copy(c, s) // (dst, src)

	fmt.Println("cpy", c)

	l := s[2:5] // [start, end)
	fmt.Println("sll:", l)

	l = s[2:] // [2, end]
	fmt.Println("sl2:", l)

	l = s[:5] // [start,end) OR the first n-elements
	fmt.Println("sl3", l)

	row := 4
	twoD := make([][]int, row)
	for i := 0; i < row; i++ {
		innerLen := i
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
