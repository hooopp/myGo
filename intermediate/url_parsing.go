package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://example.com:2323/path?query=123#fragment"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("Parsed URL:", parsedURL)
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Port:", parsedURL.Port())
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Query:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)

	rawURL1, err1 := url.Parse("https://example.com/path?name=John&age=30")
	if err1 != nil {
		fmt.Println("Error parsing URL:", err1)
		return
	}
	fmt.Println("Parsed URL 1:", rawURL1)
	queryParams := rawURL1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	values := url.Values{}

	values.Add("name", "John")
	values.Add("age", "30")

	encodedQuery := values.Encode()
	fmt.Println("Encoded Query:", encodedQuery)

	baseURL := "https://example"

	fullURL := baseURL + "?" + encodedQuery
	fmt.Println("Full URL:", fullURL)
	
}