package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type sites []string

func requestSite(req *http.Request) (resp *http.Response, err error) {
	client := http.Client{}
	res, err := client.Do(req)

	fmt.Printf("Completed req for %v\n", req.URL)

	return res, err
}

func processResponse(ctx context.Context, wg *sync.WaitGroup, path string) *http.Response {
	defer wg.Done()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	res, err := requestSite(req)

	if err != nil {
		log.Fatalf("An error occurred: %v\n", err)
	}

	return res
}

func main() {
	s := sites{"http://google.com", "http://yahoo.com", "http://bing.com"}

	var wg sync.WaitGroup
	wg.Add(len(s) - 1)

	var rs *http.Response

	for i := 0; i < len(s); i++ {
		go func(path string) {
			rs = processResponse(context.Background(), &wg, path)
		}(s[i])
	}

	wg.Wait()
	fmt.Printf("Successful res: %v\n\tStatus code: %v\n", rs.Request.URL, rs.StatusCode)

	fmt.Println("Program completed")
}
