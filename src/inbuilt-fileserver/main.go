package main

import (
  "net/http"
)

func main()  {
  // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
  //   //use the ServeFile method instead which takes a responsewriter, request and name of the file to lookup
  //   http.ServeFile(w, r, "public" + r.URL.Path)
  // }

  // you could also use the fileserver handler in place of the default serve marks. Tpo create the default file system you would have to convert string into a  Dir object.
  http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
}
