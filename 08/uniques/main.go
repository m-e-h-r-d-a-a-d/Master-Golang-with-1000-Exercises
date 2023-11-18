package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())

	var uniques []int
	max, _ := strconv.Atoi(os.Args[1])

loop:
	for i := 0; i < max; {
		n := rand.Intn(max) + 1
		for _, v := range uniques {
			if v == n {
				continue loop
			}
		}
		i++
		uniques = append(uniques, n)
	}

	fmt.Println("uniques: ", uniques)
}
