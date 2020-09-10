package main

import (
	"encoding/json"
	"fmt"
)

type (
	person struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}
)

func main() {
	p := person{
		FirstName: "Jhon",
		LastName:  "Doe",
	}

	plain, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(plain))
}

