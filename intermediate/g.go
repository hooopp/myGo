package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter file name to crate: ")
	var fileName string
	fmt.Scanln(&fileName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	file.WriteString("package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\t\n\n}\n")
}