package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) != -1
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, fn func(string) bool) bool {
	for _, v := range vs {
		if !fn(v) {
			return false
		}
	}
	return true
}

func Map(vs []string, fn func(string) string) []string {
	results := make([]string, 0)
	for _, v := range vs {
		results = append(results, fn(v))
	}
	return results
}

func Filter(vs []string, fn func(string) bool) []string {
	ret := make([]string, 0)
	for _, v := range vs {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func main() {
	strs := []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Include(strs, "peach"))
	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(arg string) bool {
		return strings.HasPrefix(arg, "p")
	}))

	fmt.Println(All(strs, func(arg string) bool {
		return strings.HasPrefix(arg, "p") || strings.HasPrefix(arg, "a")
	}))

	fmt.Println(Filter(strs, func(arg string) bool {
		return strings.HasPrefix(arg, "p")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}
