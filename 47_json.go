package main

import (
	"encoding/json"
	"fmt"
	//	"os"
)

type responsel struct {
	Page   int
	Fruits []string
}

type responsel2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	reslD := &responsel{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	reslB, _ := json.Marshal(reslD)
	fmt.Println(string(reslB))

	res2D := responsel2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num": 6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// cast
	num := dat["num"].(float64)
	fmt.Println(num)

	//	str := `{"page":1,"fruits": ["apple","peach"]}`
}
