package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// rizky := person{firstName: "Rizky", lastName: "Haksono"}
	// fmt.Println(rizky)

	var rizky person

	rizky.firstName = "Rizky"
	rizky.lastName = "Haksono"

	fmt.Println(rizky)
	fmt.Printf("%+v", rizky)
}