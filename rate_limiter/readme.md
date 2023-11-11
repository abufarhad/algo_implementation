# Rate Limiter
- The token bucket algorithm works by allocating a fixed number of tokens to each user, which they can use to make requests
- As tokens are used up, they are gradually replenished over time.
- If a user runs out of tokens, they must wait for more tokens to become available before they can make additional requests
- This approach ensures that high-traffic users are limited in the number of requests they can make.


## To run the code use
```go
go run main.go
```

## Output
![Alt text](images/output.gif)