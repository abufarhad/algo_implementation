package main

import (
	"log"
	"net/http"
)

type site struct {
	Url string
}

type result struct {
	Url    string
	Status int
}

func doJobs(wId int, jobs <-chan site, results chan<- result) {
	for site := range jobs {
		log.Println("Worker Id ", wId)
		resp, err := http.Get(site.Url)
		if err != nil {
			log.Println("error ", err.Error())
		}
		results <- result{site.Url, resp.StatusCode}
	}
}

func main() {
	numOfJobs := 3

	jobLists := []string{
		"https://facebook.com",
		"https://example.com",
		"https://google.com",
	}
	jobs := make(chan site, numOfJobs)
	results := make(chan result, numOfJobs)

	for w := 1; w <= numOfJobs; w++ {
		go doJobs(w, jobs, results)
	}

	for _, job := range jobLists {
		jobs <- site{Url: job}
	}
	close(jobs)

	for r := 1; r <= numOfJobs; r++ {
		res := <-results
		log.Println(res)
	}
}
