package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, World!")
	fmt.Println("Original data:", string(data))
	encode := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded:", encode)
	decode, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Decoded:", string(decode))

	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe Encoded:", urlSafeEncoded)
	urlSafeDecoded, err := base64.URLEncoding.DecodeString(urlSafeEncoded)
	fmt.Println("URL Safe Decoded:", string(urlSafeDecoded))
	
}

