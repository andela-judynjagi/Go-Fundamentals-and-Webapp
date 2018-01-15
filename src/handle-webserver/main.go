package main

import (
  "fmt"
  "net/http"
)

//create a handler
type myHandler struct {
  greeting string
}

//implement the handler interface by registering a method on myHandler type called ServeHTTP

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
   // a response that the handler responds to when a request comes in
   w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}

func main()  {
  //use the handle function, that takes a pattern and an instance of a handler ,to register this
  http.Handle("/", &myHandler{greeting: "Hello"})
  http.ListenAndServe(":8000", nil) // go interprets that it needs to use the default serve marks

}
