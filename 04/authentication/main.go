package main

import (
	"fmt"
	"os"
)

const (
	usage      = "Usage: [username] [password]"
	wrongUser  = "Access denied for \"%v\".\n"
	wrongPass  = "Invalid password for \"%v\".\n"
	successful = "Access granted to \"jack\"."
	user1      = "jack"
	user2      = "inanc"
	pass1      = "1888"
	pass2      = "1879"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(usage)
		return
	}

	u := os.Args[1]
	p := os.Args[2]

	if u != user1 && u != user2 {
		fmt.Printf(wrongUser, u)
		return
	}

	if (u == user1 && p != pass1) || (u == user2 && p != pass2) {
		fmt.Printf(wrongPass, u)
		return
	}

	fmt.Println(successful)
}
