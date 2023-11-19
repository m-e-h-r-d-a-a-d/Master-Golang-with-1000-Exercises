package main

import "fmt"

func main() {
	dict := map[string]string{
		"good": "خوب",
		"bad":  "بد",
		"ugly": "زشت",
	}

	for k, v := range dict {
		fmt.Printf("%q means %#v\n", k, v)
	}
}
