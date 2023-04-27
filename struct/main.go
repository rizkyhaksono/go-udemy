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

	// fmt.Println(rizky)
	rizkyPointer := &rizky
	rizkyPointer.print()
	rizkyPointer.updateName("Natee")
	rizkyPointer.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	// fmt.Printf("%+v", 	p)
	fmt.Println(p)
}