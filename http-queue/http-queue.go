package main

import (
  "fmt"
  "log"
  q "http-queue/queue"
  "net/http"
  "strconv"
)

var queue = q.NewQueue(1000)

func handleEnqueue(writter http.ResponseWriter, request *http.Request) {
  if request.Method != "POST" {
    fmt.Fprintf(writter, "This endpoint only accepts POST requests")
    return
  }

  value, error := strconv.Atoi(request.URL.Path[len("/enqueue/"):])
  
  if error != nil {
    fmt.Fprintf(writter, "Invalid value '%s'", request.URL.Path[len("/enqueue/"):])
    return
  }

  queue.Enqueue(value)
  
  fmt.Fprintf(writter, "Value %v enqueued successfully", value) 
}

func handleDequeue(writter http.ResponseWriter, request *http.Request) {
  if request.Method != "GET" {
    fmt.Fprintf(writter, "This endpoint only accepts GET requests")
    return
  }
  
  value, error := queue.Dequeue()

  if error != nil {
    fmt.Fprint(writter, error)
    return
  }

  fmt.Fprintf(writter, "Value %v dequeued successfully", value)
}

func main() {
  http.HandleFunc("/enqueue/", handleEnqueue)
  http.HandleFunc("/dequeue/", handleDequeue)

  fmt.Println("Server is running at http://localhost:8000")  
  log.Fatal(http.ListenAndServe(":8000", nil))
}
