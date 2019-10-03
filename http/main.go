package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://randomuser.me/api")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("Error2: ", err2)
		os.Exit(1)
	}
	fmt.Println(string(body))
}
