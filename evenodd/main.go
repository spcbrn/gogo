package main

import (
	"fmt"
)

func main() {
	// create a slice of ints 0 through 10
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// iterate through the slice and print out whether it's "even" or "odd"
	for _, int := range ints {

		// for zero, remark snarkily that it is not divisible
		if int == 0 {
			fmt.Println(int, "is indivisible")
			continue
		}

		// if remainder from division by 2 exists, print "odd", else print "even"
		if int%2 == 0 {
			fmt.Println(int, "is even")
		} else {
			fmt.Println(int, "is odd")
		}

	}
}
