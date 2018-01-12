package main

import (
  "fmt"
  "runtime"
  "reflect"
)

var (
  name, course string
  module float64
)

// Go can infer types as long as they are initialized eg it knows name is of type string
var (
  fullname = "Judy Njagi"
  mycourse = "Learning Go"
  mymodule = 13.60
)

// Go is strongly typed eg we can't add a +b directly. Instead you type to convert eg a into type of int then add
var (
  a = 13.600
  b = 2
)

func main(){
  fmt.Println("Hello from", runtime.GOOS)
  fmt.Println("My name is", name, "and is of type", reflect.TypeOf(name))
  fmt.Println("My name is", fullname, "and is of type", reflect.TypeOf(fullname))
  fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
  c := int(a) + b
  fmt.Println("a+b = ", c)
}
