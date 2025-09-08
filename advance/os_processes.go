package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	cmd := exec.Command("echo", "Hello, World!")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
	fmt.Println(string(output))

	cmd = exec.Command("printenv", "SHELL")
	output, err = cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
	fmt.Println(string(output))

	fmt.Println("Output:", string(output))

	pr, pw := io.Pipe()
	cmd = exec.Command("grep", "foo")
	cmd.Stdin = pr
	go func(){
		defer pw.Close()
		pw.Write([]byte("foo\nbar\nbaz\nfoo"))
	}()

	output, err = cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
	fmt.Println(string(output))

	cmd = exec.Command("ls", "-l")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(output))

	cmd = exec.Command("grep", "foo")
	cmd.Stdin = strings.NewReader("foo\nbar\nbaz\nfoo")
	output, err = cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
	fmt.Println(string(output))

	cmd = exec.Command("sleep", "5")
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting:", err)
		return
	}
	fmt.Println("Process is complete")

	cmd = exec.Command("sleep", "5")
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	time.Sleep(2 * time.Second)
	err = cmd.Process.Kill()

	if err != nil {
		fmt.Println("Error killing process:", err)
		return
	}
	fmt.Println("Process is complete")

}
