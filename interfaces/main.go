package main

import (
	"fmt"
	"reflect"
)

// will look for types that have a getGreeting func that returns a string
// and make them 'bot's
type bot interface {
	getProp(string) string
}
type langBot struct {
	language string
	greeting string
	planet   string
}
type alienBot struct {
	language string
	greeting string
	planet   string
}

func main() {
	eb := newLangBot("English", "Hello there")
	sb := newLangBot("Spanish", "Hola")
	ab := alienBot{
		language: "mOXI459",
		greeting: "GorbuluGONshUpe",
		planet:   "ZOrpaMula",
	}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(ab)
}

func newLangBot(l string, g string) langBot {
	return langBot{
		language: l,
		greeting: g,
		planet:   "Earth",
	}
}
func (lb langBot) getProp(key string) string {
	return reflect.ValueOf(lb).FieldByName(key).String()
}
func (ab alienBot) getProp(key string) string {
	return reflect.ValueOf(ab).FieldByName(key).String()
}
func printGreeting(b bot) {
	fmt.Println("In", b.getProp("language"), ":", b.getProp("greeting"), "from", b.getProp("planet"), "!")
}
