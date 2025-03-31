
package main
import "fmt"
import "time"

// A simple function to be executed as a goroutine
func printMessage(message string) {
    time.Sleep(2 * time.Second)  // 
	// Simulate some work
    fmt.Println(message)
}

func main() {
    go printMessage("Hello from Goroutine")  // This runs concurrently
    fmt.Println("Hello from main")
    time.Sleep(3 * time.Second)  // Wait for goroutine to finish
}

