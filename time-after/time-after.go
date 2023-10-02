package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	for i := 0; i < 10; i++ {
		go doRequest(ctx)
	}

	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second * 2)
}

func doRequest(ctx context.Context) {
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("timer timeout")
	case <-ctx.Done():
		fmt.Println("context cancelled")
	}
}

func doRequestValid(ctx context.Context) {
	delay := time.NewTimer(time.Second * 2)

	select {
	case <-delay.C:
		// do something after one second.
	case <-ctx.Done():
		// do something when context is finished and stop the timer.
		if !delay.Stop() {
			// if the timer has been stopped then read from the channel.
			<-delay.C
		}
	}
}
