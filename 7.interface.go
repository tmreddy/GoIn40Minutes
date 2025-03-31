package main
import "fmt"

// Define an interface
type Speaker interface {
    Speak() string
}

// Implement the interface with a struct
type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}

func main() {
    p := Person{Name: "Alice"}
    var s Speaker = p // Person satisfies the Speaker interface
    fmt.Println(s.Speak())
}
