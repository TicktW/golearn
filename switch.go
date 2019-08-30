package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "name: " + p.name + " age: " + strconv.Itoa(p.age)
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Bob", 70}
	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int , its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is an string, its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is an person, its value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is a different type, its value is %s\n", index, value)
		}
	}
}
