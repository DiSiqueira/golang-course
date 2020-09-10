package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type (
	item struct {
		Name string `json:"name"`
		SKU int `json:"sku"`
	}

	itemList struct {
		Items []item `json:"items"`
	}
)

func main() {
	fmt.Println("reading csv")
	file, err := os.Open("./test.csv")
	if err != nil {
		log.Fatalf("can not open file: %s", err)
	}

	reader := csv.NewReader(file)

	itemList := itemList{
		Items: make([]item, 0),
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("can not read csv line %s", err)
		}

		sku, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalf("can not convert %s to int", line[1])
		}

		i := item{
			Name: line[0],
			SKU:  sku,
		}

		itemList.Items = append(itemList.Items, i)

		//plain, err := json.Marshal(i)
		//if err != nil {
		//	log.Fatalf("error converting to json: %s", err)
		//}
		//
		////fmt.Printf("|%10s|%10s|\n", line[1], line[0])
		//
		//fmt.Println(string(plain))
	}

	plain, err := json.Marshal(itemList)
	if err != nil {
			log.Fatalf("error converting to json: %s", err)
	}

	fmt.Println(string(plain))

	fmt.Println("all done")
}
