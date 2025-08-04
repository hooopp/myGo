package main

import (
	"fmt"
	"path/filepath" // Assuming file_path.go is in the same package
)

func main() {
	joinedPath := filepath.Join("home", "Documents", "downloads", "file.zip")
	fmt.Println("Joined path:", joinedPath)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized path:", normalizedPath)

	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("Directory:", dir)
	fmt.Println("File:", file)
	fmt.Println(filepath.Base("/home/user/docs"))

}
