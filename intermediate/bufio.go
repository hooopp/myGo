package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(strings.NewReader("Hello bufio package!This is New!\n123213"))

	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading line:", err)
		return
	}
	fmt.Println("Read line:", line)

	writer := bufio.NewWriter(os.Stdout)
	data = []byte("Hello, bufio writer!")
	n, err = writer.Write(data)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return
	}
	fmt.Printf("Wrote string: %s (%d bytes)\n", data, n)

	fmt.Println("Flushing writer to stdout...")
	err = writer.Flush()
	fmt.Println()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
	fmt.Println("Data flushed to stdout successfully")

	str := "This is string."
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote string: %s (%d bytes)\n", str, n)

	fmt.Println("Flushing writer to stdout...")
	err = writer.Flush()
	fmt.Println()
	if err != nil {
		fmt.Println("Error flushing writer after writing string:", err)
		return
	}
}