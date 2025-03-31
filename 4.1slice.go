package main
import "fmt"

func main() {
   // Slices are more commonly used than arrays
   var slice []int = []int{1, 2, 3, 4, 5}

   // Slicing a slice (i.e., taking a subset)
   subSlice := slice[1:4]  //This will give you elements at indices 1, 2, 3
   fmt.Println(subSlice)
}
