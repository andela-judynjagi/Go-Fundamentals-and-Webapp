package main

import (
  "fmt"
  "os"
  "html/template"
)

func main()  {
  templateString := `Lemonade Stand Supply` //This's the string that forms the basis of our template. Use backticks
  t, err := template.New("title").Parse(templateString)
  // if templates fails
  if err != nil {
    fmt.Println(err)
  }
  // if execution of the template fails
  err = t.Execute(os.Stdout, nil)
  if err != nil {
    fmt.Println(err)
  }
}
