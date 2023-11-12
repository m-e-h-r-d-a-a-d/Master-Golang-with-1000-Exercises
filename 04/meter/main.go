package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no input.")
		return
	}

	feet, _ := strconv.ParseFloat(os.Args[1], 64)
	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters \n", feet, meters)
}
