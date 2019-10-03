package main

import (
	"fmt"
	"reflect"
)

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
		lastName:  "Lem",
		age:       30,
		contactInfo: contactInfo{
			email: "spcbrn1@gmail.com",
			zip:   84103,
		},
	}
	// get pointer/memory address for struct
	// mePointer := &me
	me.updateName("", "Lemke")

	// me.print()
}

func (p *person) updateName(first string, last string) {
	if first != "" {
		p.firstName = first
	}
	if last != "" {
		p.lastName = last
	}
}

func idcSlice(s reflect.Value) []int {
	o := []int{}
	it := s.NumField()
	for i := 0; i < it; i++ {
		o = append(o, i)
	}
	return o
}

func (p person) values() []reflect.Value {
	o := []reflect.Value{}
	pVal := reflect.ValueOf(p)
	idc := idcSlice(pVal)
	for i := range idc {
		o = append(o, pVal.Field(i))
	}
	return o
}

func (p person) keys() []string {
	o := []string{}
	pVal := reflect.ValueOf(p)
	idc := idcSlice(pVal)
	for i := range idc {
		fmt.Println(pVal.Field(i))
	}
	return o
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
