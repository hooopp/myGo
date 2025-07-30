package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, Go Standard Library")
	resp, err := http.Get("http://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("HTTP Response Status:", resp.Status)
}