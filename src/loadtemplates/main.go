package main

import (
  "html/template"
  "net/http"
  "log"
)

func main()  {
  templates := populateTemplates()
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
    requestedFile := r.URL.Path[1:]
    t := templates.Lookup(requestedFile + ".html")
    if t != nil {
      err := t.Execute(w, nil)
      if err != nil {
        log.Println(err)
      }
    } else {
      w.WriteHeader(http.StatusNotFound)
    }
  })
  http.Handle("/img/", http.FileServer(http.Dir("public")))
  http.Handle("/css/", http.FileServer(http.Dir("public")))
  http.ListenAndServe(":8000", nil)
}

func populateTemplates() *template.Template {
  result := template.New("template")
  const basePath = "templates"
  template.Must(result.ParseGlob(basePath + "/*.html")) // we are pulling all the packages that end in .html extention
  // Must ensures that if the template is not available then an error is thrown
  return result

}
