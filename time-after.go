package main

import (
	"context"
	"time"
)

func doRequest(ctx context.Context) {
	select {
	case <-time.After(time.Second):
		// do something after 1 second.
	case <-ctx.Done():
		// do something when context is finished.
		// resources created by the time.After() will not be garbage collected
	}
}

func doRequestValid(ctx context.Context) {
	delay := time.NewTimer(time.Second)

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
