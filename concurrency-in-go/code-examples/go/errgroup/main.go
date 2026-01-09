package main

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

var services []string = []string{
	"https://httpbin.org/delay/2", // 2 second delay
	"https://httpbin.org/delay/2",
	"https://httpbin.org/delay/2",
	"https://httpbin.org/status/503", // Immediately returns service unavailable
	"https://httpbin.org/delay/2",
	"https://httpbin.org/delay/2",
	"https://httpbin.org/delay/2",
	"https://httpbin.org/delay/2",
	"https://httpbin.org/delay/2",
	"https://httpbin.org/delay/2",
	"https://httpbin.org/delay/2",
}

var numStarted atomic.Uint32
var numHealthy atomic.Uint32

func main() {
	var ctx context.Context = context.Background()
	err := CheckSystemHealth(ctx)

	fmt.Printf("err: %v\n", err)
	fmt.Printf("Queries started: %d\n", numStarted.Load())
	fmt.Printf("Healthy: %d\n", numHealthy.Load())
}

func CheckSystemHealth(ctx context.Context) error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	errg, ctx := errgroup.WithContext(ctx)
	// For the sake of resources, only open up to 3 connections at any given time.
	errg.SetLimit(3)

	// For each service, check if it is responding.
	for idx, svc := range services {
		idx := idx
		svc := svc

		errg.Go(func() error {
			numStarted.Add(1)
			fmt.Printf("Query #%d: %s\n\n", idx, svc)
			req, _ := http.NewRequestWithContext(ctx, "GET", svc, nil)
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("%s unreachable: %w", svc, err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				return fmt.Errorf("%s returned %d", svc, resp.StatusCode)
			}

			numHealthy.Add(1)
			fmt.Printf("Success #%d: %s\n\n", idx, svc)
			return nil
		})
	}

	// If ANY service fails, stop checking the rest.
	return errg.Wait()
}
