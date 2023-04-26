package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	rizky := person{
		firstName: "Rizky",
		lastName:  "Haksono",
		contact: contactInfo{
			email:   "mrizkyhaksono",
			zipCode: 12345,
		},
	}

	fmt.Println(rizky)
}