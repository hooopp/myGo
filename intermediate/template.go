package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {

	tmpl, err := template.New("example").Parse("Welcome to {{.name}}!")
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"name": "John",
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	tmpl = template.Must(template.New("example").Parse("Welcome to {{.name}}!"))

	data = map[string]interface{}{
		"name": "Jane",
	}
	err = tmpl.Execute(os.Stdout, data)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)

	templates := map[string]string{
		"welcome": "Welcome to {{.name}}!",
		"notification": "{{.name}} has new notifications: {{.notification}}",
		"error": "An error occurred: {{.errorMessage}}",
	}

	parseTemplates := make(map[string]*template.Template)

	for name, tmpl := range templates {
		parseTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		fmt.Println("\nMenu")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notifications")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch choice {
			case "1":
				data = map[string]interface{}{"name": name}
				tmpl = parseTemplates["welcome"]
			case "2":
				fmt.Println("Enter your notification: ")
				notification, _ := reader.ReadString('\n')
				notification = strings.TrimSpace(notification)
				tmpl = parseTemplates["notification"]
				data = map[string]interface{}{"name": name, "notification": notification}
			case "3":
				fmt.Println("Enter error message: ")
				errorMessage, _ := reader.ReadString('\n')
				errorMessage = strings.TrimSpace(errorMessage)
				tmpl = parseTemplates["error"]
				data = map[string]interface{}{"name": name, "errorMessage": errorMessage}
			case "4":
				fmt.Println("Exiting...")
				return
			default:
				fmt.Println("Invalid choice, please try again.")
				continue
				
		}
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template:", err)
			continue
		}
	}
}