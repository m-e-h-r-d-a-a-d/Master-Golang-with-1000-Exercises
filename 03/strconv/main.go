package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	feetInMeters = 0.3048
	feetInYards  = feetInMeters / 0.9144
)

func main() {
	feet, _ := strconv.ParseFloat(os.Args[1], 64)
	meter := math.Round(feet * feetInMeters)
	yards := math.Round(feet * feetInYards)
	fmt.Printf("%g feet is %g meters.\n", feet, meter)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)
}
