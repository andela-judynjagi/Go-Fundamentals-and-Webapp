/*
Package main is a package declaration.
It tells the compiler to compile this program as a stand alone executable and not as a shell library
*/
package main

import (
  "fmt" //its a part of the standard Library
  "runtime"
  "reflect"
)

var (
  name, course string
  module float64
)

// Go can infer types as long as they are initialized eg it knows name is of type string
// this variables are called package level variables and they are global
var (
  fullname = "Judy Njagi"
  mycourse = "Learning Go"
  mymodule = 13.60
)

// Go is strongly typed eg we can't add a + b directly. Instead you need to convert eg a into type of int then add
var (
  a = 13.600
  b = 2
)

/*
func main is the entry point of the executable program
it takes no arguments
it doesn't return any value
*/
func main(){
  fmt.Println("Hello from", runtime.GOOS)
  fmt.Println("My name is", name, "and is of type", reflect.TypeOf(name))
  fmt.Println("My name is", fullname, "and is of type", reflect.TypeOf(fullname))
  fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
  c := int(a) + b // its short assignment and only works withing functions. Variables declared at function level must ALL be used

  /*
  Go passes arguments by value and not reference
  This means that it creates a copy of that variable via &variablename
  That copy gets a new memory address and to get its value use *variablename
  */
  age := 20

  myage := &age // this prints the memory address of variable age

  fmt.Println("the memory address of *age* is", myage, "and the value of *age* is ", *myage, "years")
  fmt.Println("a+b = ", c)

  changeAge(&age)
  fmt.Println("My new age is", age)
}


// use functions to change a variable value. Passing argumennts by reference

func changeAge(age *int) int { // first asterics tells Go that age is a pointer to an integer
  *age = 30

  fmt.Println("Trying to change my age to", *age)

  return *age
}

// contants are declared with the key word const and are immutable eg const age = 20 . This value will never change
