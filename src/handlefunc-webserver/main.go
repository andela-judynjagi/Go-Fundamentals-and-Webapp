package main

import (
  "net/http"
)

func main()  {
  /*
  a function that responds to our web requests and start up the server.
  handlefunc takes two params. First the root path we want to listen on. forward slash is for handling all requests
  Second is a call back func when a request comes on that path
  ResponseWriter writes responses back to the requester and the other params is a pointer to the request object
  */
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //lets write a response as a slice of byte
    w.Write([]byte("Hello World. Using the inbuild handleFuch method"))
  })
  // setup the server
  http.ListenAndServe(":8000", nil)
}
