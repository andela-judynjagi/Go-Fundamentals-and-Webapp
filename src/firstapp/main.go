package main

import (
  "fmt"
  "runtime"
)

var (
  name, course string
  module float64
)

func main(){
  fmt.Println("Hello from", runtime.GOOS)
  fmt.Println("My name is", name)
  fmt.Println("Module is set to", module)
}
