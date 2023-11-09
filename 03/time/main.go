package main

import (
	"fmt"
	"time"
)

func main() {
	const n = 10000
	t := time.Second * n
	fmt.Println(t)
}
