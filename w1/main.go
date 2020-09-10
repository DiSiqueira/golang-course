package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("reading csv")
	file, err := os.Open("./test.csv")
	if err != nil {
		log.Fatalf("can not open file: %s", err)
	}

	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("can not read csv line %s", err)
		}

		fmt.Printf("|%10s|%10s|\n", line[1], line[0])
	}

	fmt.Println("all done")
}
