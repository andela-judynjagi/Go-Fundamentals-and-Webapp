package main

import (
  "fmt"
)

func main(){
  // map takes a key and a value. Key has to inside brackets

  tvShow := make(map[string]string)
  tvShow["How I met your mother"] = "season1"
  fmt.Println(tvShow)

  //shorthannd formatting
  books := map[string]int {
    "suspect": 1,
    "again": 3,
  }
  fmt.Println(books)

  // iterating through a map
  for key, value := range books {
    fmt.Println("They key is:", key, "The value is:",value)
  }

  //delete an item on the map use delete
  delete(books, "again")
  fmt.Println(books)
}
