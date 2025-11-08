package main

// see https://www.calhoun.io/intro-to-templates-p4-v-in-mvc/

import (
  "net/http"
  "github.com/RodrigoMattosoSilveira/rstpl/views"
)

var index *views.View
var contact *views.View

func main() {
  index = views.NewView("bootstrap", "views/index.html")
  contact = views.NewView("bootstrap", "views/contact.html")

  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/contact", contactHandler)
  http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  index.Render(w, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
  contact.Render(w, nil)
}
