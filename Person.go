package main

import "fmt"

type Person struct {
	firstName   string
	lastName    string
	contactInfo ContactInfo
}

func (p *Person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
