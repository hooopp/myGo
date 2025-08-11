package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"os"
)

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		fmt.Println("Error closing resource:", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer
	buf.WriteString("Hello, Buffer!")
	fmt.Println(buf.String())
}

func multiReadersExample() {
	r1 := strings.NewReader("Reader 1")
	r2 := strings.NewReader("Reader 2")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) // heap
	_, err := buf.ReadFrom(mr)
	if err != nil {
		fmt.Println("Error reading from multi-reader:", err)
		return
	}
	fmt.Println("Combined output:", buf.String())
}

func readFromReader(r io.Reader){
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		fmt.Println("Error reading from reader:", err)
		return
	}
	fmt.Printf(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		fmt.Println("Error writing to writer:", err)
	}
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello from Pipe!"))
		pw.Close()
	}()
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(pr)
	if err != nil {
		fmt.Println("Error reading from pipe:", err)
		return
	}
	fmt.Println("Pipe output:", buf.String())
}

func writeToFile(filepath string, data string){
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer closeResource(file)

	writer := io.Writer(file)
	_, err = writer.Write([]byte(data))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func main() {
	fmt.Println("=== Read from Reader ===")
	readFromReader(strings.NewReader("Hello, Reader!"))

	fmt.Println("=== Write to Writer ===")
	writer := new(bytes.Buffer) // Create a new bytes.Buffer
	writeToWriter(writer, "Hello, Writer!")
	// or
	// var writer bytes.Buffer // Create a new bytes.Buffer
	// writeToWriter(&writer, "Hello, Writer!")

	fmt.Println("=== Buffer Example ===")
	bufferExample()

	fmt.Println("=== Multi Readers Example ===")
	multiReadersExample()

	fmt.Println("=== Pipe Example ===")
	pipeExample()

}
