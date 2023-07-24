package page 

import (
  "os"
)

type Page struct {
  Title string
  Body []byte
}

func (page *Page) Save() error {
  filename := "./data/" + page.Title + ".txt"
  
  return os.WriteFile(filename, page.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
  filename := "./data/" + title + ".txt"
  
  body, error := os.ReadFile(filename)
  if error != nil {
    return nil, error
  }

  return &Page{ Title: title, Body: body }, nil
}
