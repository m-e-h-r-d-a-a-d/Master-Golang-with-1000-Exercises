package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	username := os.Args[1]
	password := os.Args[2]

	if username != "jack" {
		fmt.Printf("Access denied for \"%v\".\n", username)
		return
	}

	if password != "1888" {
		fmt.Printf("Invalid password for \"%v\".\n", username)
		return
	}

	fmt.Println("Access granted to \"jack\".")
}
