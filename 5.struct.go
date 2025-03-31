package main
import "fmt"

// Define a struct
type Person struct {
    Name string
    Age  int
}

func main() {
    // Initialize a struct
    p := Person{Name: "John", Age: 30}
    fmt.Println(p)
    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)
}
