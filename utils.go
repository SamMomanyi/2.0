package main

import "fmt"

// Print a random greeting
func RandomGreeting() {
	greetings := []string{
		"Hello!",
		"Hi there!",
		"Welcome to Go!",
		"Have a great day!",
	}
	for _, g := range greetings {
		fmt.Println(g)
	}
}
