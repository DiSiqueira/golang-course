package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

type (
	item struct {
		Name string `json:"name"`
		SKU int
	}

	itemList struct {
		Items []item `json:"items"`
	}

	job struct {
		line []string
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

	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}

	jobs := Producer(reader, wg)

	for i := 0; i<3; i++ {
		go func(i int) {
			for job := range jobs {
				item := worker(i , job)

				mtx.Lock()
				itemList.Items = append(itemList.Items, item)
				mtx.Unlock()
				wg.Done()
			}
		}(i)
	}

	wg.Wait()

	plain, err := json.Marshal(itemList)
	if err != nil {
		log.Fatalf("error converting to json: %s", err)
	}

	fmt.Println(string(plain))

	fmt.Println("all done")
}

func Producer(reader *csv.Reader, wg *sync.WaitGroup) chan job {
	jobs := make(chan job, 10)
	for {
		csvLine, err := reader.Read()
		if err == io.EOF {
			close(jobs)
			break
		}
		if err != nil {
			log.Fatalf("can not read csv line %s", err)
		}

		job := job{
			line: csvLine,
		}

		jobs <- job
		wg.Add(1)
	}

	return jobs
}

func worker(number int, job job) item {
	fmt.Println(number, job.line[0], job.line[1])

	sku, err := strconv.Atoi(job.line[1])
	if err != nil {
		log.Fatal(err)
	}

	item := item{
		Name: job.line[0],
		SKU:  sku,
	}

	return item
}