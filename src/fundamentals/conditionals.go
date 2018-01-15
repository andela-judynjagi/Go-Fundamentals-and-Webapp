package main

import (
  "fmt"
  "math/rand"
  "time"
  )

func main(){
  course1 := 40
  course2 :=20

  if course1 < course2 {
    fmt.Println("Course2 is doing better than course1")
  } else if course2 < course1 {
    fmt.Println("Course1 is doing better than course2")
  } else {
    fmt.Println("Both courses are doing great")
  }

  // using simple initiliazing statement
  if price1, price2 := 50, 100; price1 < price2 {
    fmt.Println("Price2 is doing better than price1")
    if price2 <= 100 {
      fmt.Println("This is called nesting")
    }
  } else if course2 < course1 {
    fmt.Println("Price1 is doing better than price2")
  } else {
    fmt.Println("Both prices are doing great")
  }

  //switch statement
  switch "docker" {
  case "lnux":
    fmt.Println("Linux courses")
  case "windows":
    fmt.Println("Windows courses")
  case "docker":
    fmt.Println("Docker courses")
    fallthrough // it's a hack to print the next line after the loop breaks at the correct statement
  default:
    fmt.Println("No courses")
  }

  switch randomnumber := random(); randomnumber {
  case 0, 2, 4, 6, 8:
    fmt.Println("The even number is ", randomnumber)
  case 1, 3, 5, 7, 9:
    fmt.Println("The odd number is ", randomnumber)
  }
}


func random() int{
  rand.Seed(time.Now().Unix())
  return rand.Intn(10)

}
