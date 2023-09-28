package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	fakeEmail := contactInfo{"fake@gmai.com", 12345}
	alex := person{firstName: "Alex", lastName: "Anderson", contactInfo: contactInfo{email: "alex@gmail.com", zip: 12345}}
	fakeEmail.print()
	alex.print()
	alex.updateName("Mark")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (c contactInfo) print() {
	fmt.Printf("%+v\n", c)
}
