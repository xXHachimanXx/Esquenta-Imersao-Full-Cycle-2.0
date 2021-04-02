package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	requestId := make(chan int)
	concurrency := 2
	for i := 1; i <= concurrency; i++ {
		worker(requestId, i)
	}

	for x := 0; x < 10; x++ {
		requestId <- x
	}
}

func worker(requestId chan int, worker int) {
	for r := range requestId {
		res, err := http.Get("http://localhost:8585/products")

		if err != nil {
			log.Fatal("Erro")
		}

		defer res.Body.Close()

		content, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("Worker %d, RequestId: %d, Content: %s", worker, r, string(content))
		time.Sleep(time.Second * 2)
	}
}
