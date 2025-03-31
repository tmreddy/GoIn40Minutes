package main

import "fmt"

func main() {
    var age int = 25           // Declaring a variable of type int
    var name string = "Alice"  // Declaring a variable of type string

    // You can also use the shorthand declaration with :=
    height := 5.7  // Implicitly declared height as a float64

    fmt.Println(name, "is", age, "years old and", height, "feet tall.")
}
