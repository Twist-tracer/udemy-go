package main

func main() {
	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}
