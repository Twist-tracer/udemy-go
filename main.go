package main

import "fmt"

func main() {
	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		contact: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)
}
