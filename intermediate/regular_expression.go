package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("He said \n\"I am great\"")
	fmt.Println("He said I am great")

	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	email := "user@gmail.com"
	email2 := "invalid_email"

	fmt.Println("Email1: ", re.MatchString(email))
	fmt.Println("Email2: ", re.MatchString(email2))

	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	date := "2023-10-01"

	subMatches := re.FindStringSubmatch(date)
	fmt.Println("Date: ", subMatches[0])
	fmt.Println("year: ", subMatches[1])
	fmt.Println("month: ", subMatches[2])
	fmt.Println("day: ", subMatches[3])
}