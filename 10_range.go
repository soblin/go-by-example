package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("3 is index ", i)
		}
	}

	kvs := map[string]int{"key1": 10, "key2": 20, "key3": 30}
	for k, v := range kvs {
		fmt.Println(k, "=>", v)
	}

	for i, c := range "xyz" {
		fmt.Println(i, c)
	}
}
