package main

import (
	"fmt"
	"os"
)

func main(){
	file, error := os.Create("output.txt")
	if error != nil {
		fmt.Println("Error writing to file:", err)
	}
	defer file.Close()

	data := []byte("Hello, World!\n")
	_, err := file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	fmt.Println("Data written to file successfully")

	file, err = os.Create("writeString.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	file.WriteString("Hello Go!\n")
}