package main

import (
    "fmt"
    "time"
)

func task() {
    // Simulate a long-running task
    time.Sleep(5 * time.Second)
    fmt.Println("Task completed")
}

func main() {
    // Create a channel to communicate the completion of the task
    done := make(chan bool)

    // Start the task in a goroutine
    go func() {
        task()
        done <- true
    }()

    // Wait for the task to finish or for the timeout to occur
    select {
    case <-done:
        fmt.Println("Task completed within 3 seconds")
    case <-time.After(3 * time.Second):
        fmt.Println("Timeout: Task did not complete within 3 seconds")
    }
}
