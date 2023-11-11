// Least Connection algorithm
package main

import (
	"load_balancer/config"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type Server struct {
	URL               *url.URL   // URL of the backend server.
	ActiveConnections int        // Count of active connections
	Mutex             sync.Mutex // A mutex for safe concurrency
	Healthy           bool
}

func nextServerLeastActive(servers []*Server) *Server {
	leastActiveConnections := -1
	leastActiveServer := servers[0]
	for _, server := range servers {
		server.Mutex.Lock()
		if (server.ActiveConnections < leastActiveConnections || leastActiveConnections == -1) && server.Healthy {
			leastActiveConnections = server.ActiveConnections
			leastActiveServer = server
		}
		server.Mutex.Unlock()
	}

	return leastActiveServer
}

func (s *Server) Proxy() *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(s.URL)
}

func main() {
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err.Error())
	}

	healthCheckInterval, err := time.ParseDuration(config.HealthCheckInterval)
	if err != nil {
		log.Fatalf("Invalid health check interval: %s", err.Error())
	}

	var servers []*Server
	for _, serverUrl := range config.Servers {
		u, _ := url.Parse(serverUrl)
		servers = append(servers, &Server{URL: u})
	}

	for _, server := range servers {
		go func(s *Server) {
			for range time.Tick(healthCheckInterval) {
				res, err := http.Get(s.URL.String())
				if err != nil || res.StatusCode >= 500 {
					s.Healthy = false
				} else {
					s.Healthy = true
				}
			}
		}(server)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server := nextServerLeastActive(servers)

		server.Mutex.Lock()
		server.ActiveConnections++
		server.Mutex.Unlock()

		server.Proxy().ServeHTTP(w, r)

		server.Mutex.Lock()
		server.ActiveConnections--
		server.Mutex.Unlock()
	})

	log.Println("Starting server on port", config.ListenPort)
	err = http.ListenAndServe(config.ListenPort, nil)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
