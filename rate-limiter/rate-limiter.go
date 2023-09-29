package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
)

type Request struct {
	Payload string
}

type Client interface {
	SendRequest(ctx context.Context, request Request) error
}

type client struct {
}

func (c client) SendRequest(ctx context.Context, request Request) error {
	fmt.Println("send request", request.Payload)
	return nil
}

func main() {
	ctx := context.Background()
	c := client{}
	requests := make([]Request, 100)
	for i := 0; i < 100; i++ {
		requests[i] = Request{Payload: strconv.Itoa(i)}
	}
	log := logrus.New()
	makeBatchApiCalls(ctx, c, log, requests)
}

// TODO: rate limit api calls
// 1. In fly requests limit
// 2. Per second limit
func makeBatchApiCalls(ctx context.Context, c Client, log *logrus.Logger, requests []Request) {
	wg := sync.WaitGroup{}
	for _, r := range requests {
		r := r
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := c.SendRequest(ctx, r)
			if err != nil {
				log.WithError(err).Error("send request")
			}
		}()
	}
	wg.Wait()
}
