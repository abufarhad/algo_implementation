package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"load_balancer_rr/lb"
	"log"
	"net/http"
	"sync"
)

func handleRoot(serverName string) gin.HandlerFunc {
	fmt.Println("BackendServerName", serverName)
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"BackendServerName": serverName,
			"ResponseHeader":    c.Request.Header,
		})
	}
}

func serveBackend(serverName, port string) {
	r := gin.Default()

	// Use handleRoot function with the specified serverName
	r.GET("/", handleRoot(serverName))

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error serving backend:", err)
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(5)

	go func() {
		lb.Serve("config.json")
		wg.Done()
	}()

	go func() {
		serveBackend("google", "8081")
		wg.Done()
	}()

	go func() {
		serveBackend("bing", "8082")
		wg.Done()
	}()

	go func() {
		serveBackend("facebook", "8083")
		wg.Done()
	}()

	go func() {
		serveBackend("youtube", "8084")
		wg.Done()
	}()

	wg.Wait()
}
