package main

import "fmt"

func main() {

	// // Define map
	// emails := make(map[string]string)

	// // assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"

	// Define map and assign key values
	var emails = map[string]string{
		"Bob":    "bob@gmail.com",
		"Sharon": "sharon@gmail.com",
	}

	fmt.Println(emails)
}
