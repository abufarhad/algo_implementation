package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Step 1: Define the Task
// A task that accepts a URL and returns the extracted data as a string.
type Task func(url string) (string, error)

// Step 2: Create the Worker
// A worker is a goroutine that processes tasks and sends the results through a channel.
type Worker struct {
	id         int
	taskQueue  <-chan string
	resultChan chan<- Result
}

func (w *Worker) Start() {
	go func() {
		for url := range w.taskQueue {
			data, err := fetchAndProcess(url) // Perform the web scraping task
			w.resultChan <- Result{workerID: w.id, url: url, data: data, err: err}
		}
	}()
}

// Step 3: Implement the Worker Pool
// The worker pool manages the workers, distributes tasks, and collects results.
type WorkerPool struct {
	taskQueue   chan string
	resultChan  chan Result
	workerCount int
}

type Result struct {
	workerID int
	url      string
	data     string
	err      error
}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		taskQueue:   make(chan string),
		resultChan:  make(chan Result),
		workerCount: workerCount,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		worker := Worker{id: i, taskQueue: wp.taskQueue, resultChan: wp.resultChan}
		worker.Start()
	}
}

func (wp *WorkerPool) Submit(url string) {
	wp.taskQueue <- url
}

func (wp *WorkerPool) GetResult() Result {
	return <-wp.resultChan
}

// Fetch and process the data from the URL
func fetchAndProcess(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to fetch the URL")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Process the fetched data and extract the required information
	// I would use a library like 'goquery' to parse the HTML and extract the relevant data. You have do it yourself 🤣
	extractedData := processData(string(body))

	return extractedData, nil
}

// function to process the data, replace this with actual processing logic
func processData(body string) string {
	return body
}

func main() {
	urls := []string{
		"https://google.com",
		"https://bing.com",
		"https://apple.com",
	}

	workerPool := NewWorkerPool(3) // Create a worker pool with 3 workers
	workerPool.Start()

	// Submit the URLs to the worker pool for processing
	for _, url := range urls {
		workerPool.Submit(url)
	}

	// Collect the results and handle any errors
	for i := 0; i < len(urls); i++ {
		result := workerPool.GetResult()
		if result.err != nil {
			fmt.Printf("Worker ID: %d, URL: %s, Error: %v\n", result.workerID, result.url, result.err)
		} else {
			fmt.Printf("Worker ID: %d, URL: %s, Data: %s\n", result.workerID, result.url, result.data)
			// Save the extracted data to the database or process it further
			saveToDatabase(result.url, result.data)
		}
	}
}

// function to save the data to the database, replace this with actual database logic
func saveToDatabase(url, data string) {
	// Save the data to the database
}
