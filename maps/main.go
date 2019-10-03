package main

import (
	"fmt"
)

type stringMap map[string]string

func main() {
	colors := stringMap{
		"red":    "#eb3434",
		"yellow": "#ebdc34",
		"green":  "#34eb40",
	}
	colors.mergeIn(stringMap{
		"skyblue": "#34ebdf",
		"white":   "#ffffff",
	})
	fmt.Println(colors.keys())
	fmt.Println(colors.values())
	fmt.Println(colors.entries())
	colors.print("The hex code for", "is")
}

func (m stringMap) mergeIn(new stringMap) {
	for k, v := range new {
		m[k] = v
	}
}

func (m stringMap) keys() []string {
	o := []string{}
	for k := range m {
		o = append(o, k)
	}
	return o
}

func (m stringMap) values() []string {
	o := []string{}
	for _, v := range m {
		o = append(o, v)
	}
	return o
}

func (m stringMap) entries() [][2]string {
	o := [][2]string{}
	for k, v := range m {
		o = append(o, [2]string{k, v})
	}
	return o
}

func (m stringMap) print(kStr string, vStr string) {
	for k, v := range m {
		fmt.Println(kStr, k, vStr, v)
	}
}
