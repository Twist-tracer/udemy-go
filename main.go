package main

import "fmt"

func main() {
	//alex := person{firstName: "Alex", lastName: "Anderson"}

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
