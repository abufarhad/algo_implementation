// Implementing a rate limiter in Go using token bucket algorithm

package main

import (
	"fmt"
	"math"
	"time"
)

// TokenBucket Structure
type TokenBucket struct {
	tokens         float64
	maxTokens      float64
	refillRate     float64
	lastRefillTime time.Time
}

// Initialize the Token Bucket
func NewTokenBucket(maxTokens, refillRate float64) *TokenBucket {
	return &TokenBucket{
		tokens:         maxTokens,
		maxTokens:      maxTokens,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

func (tb *TokenBucket) Refill() {
	now := time.Now()
	duration := now.Sub(tb.lastRefillTime)
	tokensToAdd := duration.Seconds() * tb.refillRate
	tb.tokens = math.Min(tb.tokens+tokensToAdd, tb.maxTokens)
	tb.lastRefillTime = now

	fmt.Println(now, " - ", duration, " - ", tokensToAdd, " - ", tb.tokens)
}

func (tb *TokenBucket) Request(amount float64) bool {
	tb.Refill()
	if tb.tokens >= amount {
		tb.tokens -= amount
		return true
	}
	return false
}

func main() {
	tb := NewTokenBucket(10, 1)
	fmt.Println("tb", tb)
	for i := 0; i < 20; i++ {
		fmt.Printf("Request %d: %v\n", i+1, tb.Request(1))
		time.Sleep(500 * time.Millisecond)
	}
}
