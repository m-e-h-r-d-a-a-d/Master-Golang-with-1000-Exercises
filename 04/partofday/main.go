package main

import "time"

func main() {
	switch now := time.Now().Hour(); {
	case now < 12 && now > 6:
		println("good morning")
	case now < 16:
		println("good afternoon")
	case now < 18:
		println("good evening")
	default:
		println("good night")
	}
}
