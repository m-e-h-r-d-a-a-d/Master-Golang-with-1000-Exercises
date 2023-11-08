package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	feet, _ := strconv.ParseFloat(os.Args[1], 64)
	meter := feet * 0.3048
	fmt.Printf("%f_feet is %f meters.\n", feet, meter)
	fmt.Printf("%g_feet is %g meters.\n", feet, meter)
}
