package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "name: " + h.name + ", age: " + strconv.Itoa(h.age) + " years old, phone: " + h.phone
}

func main() {
	Bob := Human{"Bob", 39, "132-456789"}
	fmt.Println("This human is: ", Bob)
}
