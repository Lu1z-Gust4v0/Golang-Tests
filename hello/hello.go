package main

import (
	"fmt"
	"greetings/greetings"
	"log"
)

func main() {
  log.SetPrefix("greetings: ")
  log.SetFlags(0)
  
  var name string 
  fmt.Println("Enter your name: ");
  fmt.Scanln(&name)

  message, error := greetings.Hello(name)

  if (error != nil) {
    log.Fatal(error)
  }

  fmt.Println(message)
}
