package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func filepath(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			wg.Add(1)
			go filepath(path + "/" + file.Name())
		} else {
			fmt.Println(path + "/" + file.Name())
		}
	}
	wg.Done()
}

func main() {
	start := time.Now()
	file_path := os.Args[1]
	fmt.Println("FILE PATH IS", file_path)
	wg.Add(1)
	filepath(file_path)
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("TOTAL TIME:", elapsed)
}
