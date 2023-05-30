package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo ContactInfo
}

type ContactInfo struct {
	email   string
	address string
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateAddress(newAddress string) {
	p.contactInfo.address = newAddress
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Wazowski",
		contactInfo: ContactInfo{
			email:   "jim@example.com",
			address: "123 Main St",
		},
	}

	jim.print()

	(&jim).updateAddress("456 Elm St")

	jim.print()
}
