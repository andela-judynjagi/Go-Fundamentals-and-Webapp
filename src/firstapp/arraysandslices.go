package main

import (
  "fmt"
)

func main() {
  /* declare using the make method which take 3 args
  the type, length and capacity. Capacity is the max size of the slice
  */
  mycourses := make([]int, 5, 10)
  fmt.Println(len(mycourses), cap(mycourses))

  // use shorthannd formatting. If you don't specify the cap it's set to the length of the slice
  courses := []string {"judy", "Njagi"}
  fmt.Println(len(courses), cap(courses))

  // slice a slice
  mySlice := []int {1,2,3,4,5,6,7,8,9}
  sliceOfSlice := mySlice[2:5] // the left hand is include and right excluded. 2 through 5 is indexing so it will could index 2, 3, 4
  slice1 := mySlice[:5] // assumed 0 through 5
  slice2 := mySlice[5:] //assumed lenght of array-1 as last value
  fmt.Println(sliceOfSlice)
  fmt.Println("slice1",slice1)
  fmt.Println("slice2",slice2)

  // changing value at index
  mySlice[2] = 15
  fmt.Println(mySlice)

  //appending values to a slice. When the array is full append creates ann array twice the previous size
  for i :=1; i<17; i++ {
    mycourses = append(mycourses, i)
    fmt.Println(cap(mycourses))
  }

  //append slices to sliceOfSlice
  newSlice := []int {56, 78}
  mycourses = append(mycourses, newSlice...) // use ellipses to append a collection to the end of the slice
  fmt.Println(mycourses)
}
