package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
  "myhttp/page"
)

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func renderTemplate(w http.ResponseWriter, name string, page *page.Page) {
  error := templates.ExecuteTemplate(w, name + ".html", page) 

  if error != nil {
    http.Error(w, error.Error(), http.StatusInternalServerError)
  }
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
  page, error := page.LoadPage(title)
  
  if error != nil {
    http.Redirect(w, r, "/edit/" + title, http.StatusFound)
    return
  }

  renderTemplate(w, "view", page) 
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
  _page, error := page.LoadPage(title)

  if error != nil {
    _page = &page.Page{ Title: title, Body: []byte("") }
  }
  
  renderTemplate(w, "edit", _page)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
  body := r.FormValue("body")
  page := &page.Page{ Title: title, Body: []byte(body) }
  
  error := page.Save()

  if error != nil {
    http.Error(w, error.Error(), http.StatusInternalServerError)
    return
  }

  http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    match := validPath.FindStringSubmatch(r.URL.Path)

    if match == nil {
      http.NotFound(w, r)
      return
    }

    fn(w, r, match[2])
  }
}

func main() {
  // handle static files 
  http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
  
  http.HandleFunc("/view/", makeHandler(viewHandler))
  http.HandleFunc("/edit/", makeHandler(editHandler))
  http.HandleFunc("/save/", makeHandler(saveHandler))

  fmt.Println("Server is listening at port 8000")

  log.Fatal(http.ListenAndServe(":8000", nil))
}
