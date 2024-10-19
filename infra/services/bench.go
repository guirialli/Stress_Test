package services

import (
	"fmt"
	"github.com/guirialli/stress_test/infra/dtos"
	"net/http"
	"sync"
	"time"
)

func requestUrl(url string) (int, error) {
	client := &http.Client{Timeout: time.Second * 100}
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func BenchUrl(url string, requests, concurrency int) (*dtos.ResulStatusCode, error) {
	var wg sync.WaitGroup
	requestsChan := make(chan int)
	result := dtos.NewResulStatusCode(requests, concurrency)

	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range requestsChan {
				status, err := requestUrl(url)
				if err != nil {
					fmt.Printf("Erro: %v\n", err)
					result.IncrementErrors()
					continue
				}
				result.Increment(status)
			}
		}()
	}

	for i := 0; i < requests; i++ {
		requestsChan <- i
	}

	close(requestsChan)
	wg.Wait()

	result.SetTotalTime(time.Since(startTime))
	return result, nil
}
