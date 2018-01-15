package main

import "fmt"

func main(){
  //define a new type
  type Room struct {
    color string
    doors int
    size string
  }
  /*
  Above we have defined a new type called Room.
  But we haven't declared and initialized variable.
  To do so check below
  */
  Hotel :=new(Room) // we are creating an instancce variable which is initiliazed with zero values
  var Dojo Room // another way of doing it


  // To initiliaze wih values do

  Amity := Room {
    color : "Red",
    doors : 4,
    size:  "50 square meters",
  }
  fmt.Println(Amity)
  fmt.Println(Hotel)
  fmt.Println(Dojo)
  fmt.Println("\nThe color is: ", Amity.color)
}
// structs don't have objects, don't have class keyword and don't support inheritance
