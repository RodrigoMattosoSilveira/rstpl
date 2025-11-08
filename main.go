package main

// see https://www.calhoun.io/intro-to-templates-p4-v-in-mvc/

import (
  "html/template"
  "net/http"
  "path/filepath"
)

var LayoutDir string = "views/layouts"
// var bootstrap *template.Template
var index *template.Template
var contact *template.Template

func main() {
  var err error
//   bootstrap, err = template.ParseFiles(layoutFiles()...)
  files := append(layoutFiles(), "views/index.html")
  index, err = template.ParseFiles(files...)
  if err != nil {
    panic(err)
  }
  files = append(layoutFiles(), "views/contact.html")
  contact, err = template.ParseFiles(files...)
  if err != nil {
    panic(err)
  }

  http.HandleFunc("/", handler)
  http.HandleFunc("/contact", contactHandler)
  http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
//   bootstrap.ExecuteTemplate(w, "bootstrap", nil)
 	index.ExecuteTemplate(w, "bootstrap", nil)
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
  contact.ExecuteTemplate(w, "bootstrap", nil)
}

func layoutFiles() []string {
  files, err := filepath.Glob(LayoutDir + "/*.html")
  if err != nil {
    panic(err)
  }
  return files
}
