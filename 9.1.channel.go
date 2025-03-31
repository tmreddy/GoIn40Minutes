package main
import "fmt"

func greet(name string, ch chan string) {
    ch <- "Hello, " + name  // Send a message to the channel
}

func main() {
    ch := make(chan string)  // Create a new channel
    go greet("Alice", ch)    // Start a goroutine
    message := <-ch          // Receive the message from the channel
    fmt.Println(message)
}
