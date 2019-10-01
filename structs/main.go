package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

func main() {

	// second way
	me := person{
		firstName: "Chris",
		lastName:  "Lemke",
		age:       30,
		contactInfo: contactInfo{
			email: "spcbrn1@gmail.com",
			zip:   84103,
		},
	}

	me.printDetails()
}

func (p *person) updateName(first string, last string) {
	if first != "" {
		p.firstName = first
	}
	if last != "" {
		p.lastName = last
	}
}

func (p person) printDetails() {
	fmt.Printf("%+v", p)
}
