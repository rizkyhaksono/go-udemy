package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	rizky := person{
		firstName: "Rizky",
		lastName:  "Haksono",
		contactInfo: contactInfo{
			email:   "mrizkyhaksono",
			zipCode: 12345,
		},
	}

	// rizkyPointer := &rizky
	rizky.print()
	rizky.updateName("Natee")
	rizky.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	// fmt.Printf("%+v", 	p)
	fmt.Println(p)
}