package main

import (
	"encoding/json"
	"log"
	"net/http"

	ratelimit "github.com/faris789/ratelimiting"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	message := Message{
		Status: "Successful",
		Body:   "Hi! You've reached the API. How can I help you?",
	}
	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}
}

func rateLimiter(next func(w http.ResponseWriter, r *http.Request), lim *ratelimit.TokenBucket) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !lim.Allow() {
			message := Message{
				Status: "Request Failed",
				Body:   "The API is at full capacity, try again later.",
			}

			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return
		} else {
			next(w, r)
		}
	})
}

func main() {
	ratelimit.SetTokenRefreshInterval(60)
	limiter := ratelimit.NewTokenBucket(150)
	go limiter.DoRateLimiting()

	http.Handle("/sqs/metadata", rateLimiter(endpointHandler, limiter))
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println("There was an error listening on port :8081", err)
	}

}
