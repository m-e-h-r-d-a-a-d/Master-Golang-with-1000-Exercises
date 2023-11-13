package main

import "fmt"

func main() {
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Every Thing Ship",
		"Kafka's Revenge 2nd Edition",
	}

	fmt.Printf("book : %#v\n", books)
}
