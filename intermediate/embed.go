package main

import (
	"embed"
	"fmt"
)

//go:embed embedDir
var content string

//go:embed embedDir
var basicFolder embed.FS

func main() {
	fmt.Println("Embedded content:", content)
	content, err := basicFolder.ReadFile("embedDir/hello.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(content))

	


}
