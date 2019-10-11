package main

import (
	"fmt"
	"io"
	"os"
)

/*

- read in contents of a file on execution
- reference os.Args and os/#Open
- acknowledge what interface File type implements
- if 'Reader', reuse Copy func?

*/

type logWriter struct{}

func main() {
	fPath := os.Args[1]
	file := getFile(fPath)
	lw := logWriter{}
	io.Copy(lw, file)
}

func getFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return file
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote", len(bs), "bytes.")
	return len(bs), nil
}
