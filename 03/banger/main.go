package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	i := os.Args[1]
	l := len(i)
	s := i + strings.Repeat("!", l)
	fmt.Println(s)
	fmt.Println(strings.ToUpper(s))
}
