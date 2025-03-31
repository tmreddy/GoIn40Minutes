package main
import "fmt"

func main() {
    // Create a map
    students := map[string]int{
        "Alice": 25,
        "Bob":   30,
        "Eve":   22,
    }

    // Add a new key-value pair
    students["John"] = 40

    // Update value
    students["Alice"] = 26

    // Retrieve value
    age, exists := students["Eve"]
    if exists {
        fmt.Println("Eve's Age:", age)
    }
    // Delete an entry
    delete(students, "Bob")

    // Iterate over the map
    for name, age := range students {
        fmt.Println(name, "is", age, "years old")
    }
}
