package main

import (
	"fmt"
	"io"
	"os"
)

// type termWrite struct{}

func main() {
	fPath := os.Args[1]
	f := getFile(fPath)
	// lw := termWrite{}
	fmt.Println(fPath, "contents:")
	io.Copy(os.Stdout, f)
}

func getFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return file
}

// func (termWrite) Write(bs []byte) (int, error) {
// 	fmt.Println(string(bs))
// 	return len(bs), nil
// }
