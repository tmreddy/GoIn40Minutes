package main
import "fmt"

// Function that adds two integers and returns the result
func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(3, 4)
    fmt.Println("Sum:", sum)
	a, b := swap(1, 2)
	fmt.Println("Swapped:", a, b)
}

func swap(a, b int) (int, int) {
    return b, a
}
