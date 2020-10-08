package main

import (
	"fmt"
	"log"

	"github.com/disiqueira/golang-course/w6/internal/app"
)

func main() {
	a := app.NewApp()
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finishing service...")
}

