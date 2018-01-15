package main

import (
  "fmt"
  "time"
  "sync"
)

func main()  {
  // create a channel

  chan := make(chan string, 5)
  // a channel is a conduit for exchanging data bewteen two go routines
  var waitGrp sync.WaitGroup
  waitGrp.Add(2)

  go func(){
    defer waitGrp.Done()
    time.Sleep(5 * time.Second)
    fmt.Println("Hello")
  }()
  go func ()  {
    defer waitGrp.Done()
    fmt.Println("Second concurrent function")
  }()
  waitGrp.Wait()
}
/*
by adding go to a function you are transforming it to a go routine
adding () after a function makes it a self executing function
The second func will execute before the first because our program is executed line by line.
When it executes line 15 and only the first go routinne goes to sleep for 5 seconds and hence it exits and runs the second process
It's good to look at main as a go routine as well because when it exits all other processes exit. So sync helps to tell main when to exit
*/
