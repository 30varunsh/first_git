package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct { //structure declaration
	firstName string
	lastName  string
	contact   contactInfo // it contains email zipCode
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Anderson"} //struct variable
	//fmt.Println(alex)

	//-------

	// var alex person
	// alex.firstName = "varun"
	// alex.lastName = "sharma"
	// fmt.Printf("%+v", alex)

	//--------

	jim := person{
		firstName: "jim",
		lastName:  "party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 110032,
		},
	}
	jimPointer := &jim // store address of jim to jimPointer
	jimPointer.updateName("jimmyy")
	jim.print()
}

func (pointerToperson *person) updateName(newFirstName string) {
	(*pointerToperson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v ", p)
}
