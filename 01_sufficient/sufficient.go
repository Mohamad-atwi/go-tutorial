package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p person) print() {
	fmt.Printf("%v\n", p)
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Wazowski",
	}

	jim.print()
	// call "updateAddress" here
	jim.print()
}
