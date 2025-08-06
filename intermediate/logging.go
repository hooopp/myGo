package main

import (
	"log"
	"os"
)

func main() {
	log.Println("This is a log message.")
	log.SetPrefix("INFO: ")
	log.Println("This is another log message with a prefix.")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a log message with date, time.")

	infoLogger.Print("This is an info message.")
	warnLogger.Println("This is a warning message.")
	errorLogger.Println("This is an error message.")

}

var (
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)


