package main

import "fmt"

const bolingF = 212.0 // constant assignment

func main() {
	var f = bolingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F, or %g C \n", f, c)
}
