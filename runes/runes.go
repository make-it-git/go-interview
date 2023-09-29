package main

import (
	"fmt"
)

func main() {
	s := "Boat 🚢."
	fmt.Println("length of string?", s)

	for i, c := range s {
		fmt.Printf("Position %d of '%s'\n", i, string(c))
	}
}
