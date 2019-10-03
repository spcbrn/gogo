package main

import "fmt"

func main() {
	whocares := "bill"
	bp := &whocares
	fmt.Println(&whocares)
	printPointer(bp)
}

func printPointer(p *string) {
	fmt.Println(&p)
}
