package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}

func main() {
	users := []user{
		{"inanc", "1234", nil},
		{"god", "42", permissions{"admin": true}},
		{"devil", "GGG", permissions{"write": true}},
	}

	out1, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	out2, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(out1))
	fmt.Println(string(out2))
}
