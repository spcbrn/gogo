package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	runHTTP3()
}

// dispatch http request and return response
func runReq() *http.Response {
	resp, err := http.Get("https://randomuser.me/api/?results=30")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return resp
}

// first example using local bs and Read method
func runHTTP1() {
	resp := runReq()

	bs := make([]byte, 32*1024)
	resp.Body.Read(bs)
	bs = bytes.Trim(bs, "\x00")
	fmt.Println(string(bs))
}

// second example using Stdout as writer
func runHTTP2() {
	resp := runReq()
	// Stdout implements File which implements Writer, Body implements Reader
	io.Copy(os.Stdout, resp.Body)
}

// third example using custom Write function (logWriter as interface)
func runHTTP3() {
	resp := runReq()
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// custom Write function
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote", len(bs), "bytes.")
	return len(bs), nil
}

// experimental example/winging it
func runHTTPExp() {
	resp := runReq()

	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("Error2: ", err2)
		os.Exit(1)
	}
	fmt.Println(string(body))
}
