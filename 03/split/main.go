package main

import (
	"fmt"
	"path"
)

func main() {
	dir, file := path.Split("d:/project/sample.txt")
	fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)
}
