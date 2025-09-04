package main
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World from Golang")

	// Using functions from other files
	sum := Add(10, 20)
	fmt.Println("10 + 20 =", sum)

	reversed := Reverse("Golang")
	fmt.Println("Reversed:", reversed)

	RandomGreeting()
}
