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
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// var con contactInfo
	// con.email = "alex.anderson@a.com"
	// con.zipCode = 123
	// alex.contact = con

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	//fmt.Printf("%+v", jim)
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jimPointer.print()
}

func (pointerToperson *person) updateName(newFirstName string) {
	(*pointerToperson).firstName = newFirstName
}

func (pointerToperson *person) print() {
	fmt.Printf("%+v", *pointerToperson)
}
