package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	subcommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)
	
	firstFlagSub1 := subcommand1.Bool("processing", false, "description: Process the first subcommand")
	secondFlagSub1 := subcommand1.Int("bytes", 1024, "description: Process the second subcommand")

	firstFlagSub2 := subcommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("expected 'firstSub' or 'secondSub' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
		case "firstSub":
			subcommand1.Parse(os.Args[2:])
			fmt.Println("subCommand1:")
			fmt.Println("Processing:", *firstFlagSub1)
			fmt.Println("bytes:", *secondFlagSub1)
		case "secondSub":
			subcommand2.Parse(os.Args[2:])
			fmt.Println("subCommand2:")
			fmt.Println("Processing:", *firstFlagSub2)
		default:
			fmt.Println("expected 'firstSub' or 'secondSub' subcommands")
			os.Exit(1)
	}

}
