package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()
	fmt.Println("File opened successfully:", file.Name())
	data := make([]byte, 100) // Buffer to hold file data
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Print("File content:")
	fmt.Println(string(data))
	file, err = os.Open("output.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File read successfully")
}
