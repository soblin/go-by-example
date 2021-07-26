package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	ret := person{name: name, age: 42}
	return &ret
}

func newPerson2(name string) person {
	ret := person{name: name, age: 42}
	return ret
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{age: 30, name: "Alice"})
	fmt.Println(person{name: "Fred"}) // the rest is initialized with zero-value

	john := newPerson("John")
	fmt.Println(john.name, john.age)
}
