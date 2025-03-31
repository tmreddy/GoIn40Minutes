package main
import "fmt"

func main() {
   var x int = 58
   var p *int = &x  //Pointer to x

   fmt.Println("Value of x:", x)    // 58
   fmt.Println("Pointer to x:", p)  // Memory address
  fmt.Println("Value pointed by p:", *p)  // Dereferencing the pointer to get the value
}
