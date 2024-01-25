package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{firstName: "Jim", lastName: "Party", contact: contactInfo{email: "jim@email.com", zipCode: 75038}}
	fmt.Printf("%+v", jim)
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

}
func (p person) print() {
	fmt.Printf("%+v", p)

}
