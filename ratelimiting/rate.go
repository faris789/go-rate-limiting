package ratelimit

import (
	"fmt"
	"log"
	"time"
)

var ticker *time.Ticker
var gotReqCh chan bool
var acceptReqCh chan bool

type TokenBucket struct {
	RemainingTokenCount int
	BucketMaxSize       int
}

func NewTokenBucket(bucketMaxSize int) *TokenBucket {
	return &TokenBucket{
		RemainingTokenCount: bucketMaxSize,
		BucketMaxSize:       bucketMaxSize,
	}
}

func SetTokenRefreshInterval(tokenRefreshIntervalSeconds int) {
	ticker = time.NewTicker(time.Duration(tokenRefreshIntervalSeconds) * time.Second)
}

func (t *TokenBucket) Allow() bool {
	gotReqCh <- true
	return <-acceptReqCh
}

func init() {
	gotReqCh = make(chan bool)
	acceptReqCh = make(chan bool)
}

func (t *TokenBucket) DoRateLimiting() {
	for {
		select {

		case refilltime := <-ticker.C:
			fmt.Println(refilltime)
			if t.RemainingTokenCount != t.BucketMaxSize {
				t.RemainingTokenCount = t.BucketMaxSize
				log.Printf("Refill Triggered at %v", refilltime)
			} else {
				log.Printf("Refill Check Triggered at %v", refilltime)
			}
		case newReqBool := <-gotReqCh:
			if newReqBool {
				if t.RemainingTokenCount > 0 {
					t.RemainingTokenCount--
					log.Printf("Remaining Tokens: %v", t.RemainingTokenCount)
					acceptReqCh <- true
				} else {
					acceptReqCh <- false
				}

			}

		}
	}
}
