package main

import "fmt"
import "time"

// Define the amount of workers of each type 
var ORDERS_COUNT        int = 10 
var ATTENDANTS_COUNT    int = 1
var MACHINES_COUNT      int = 2
var BAKERS_COUNT        int = 1
var PACKERS_COUNT       int = 1
var DELIVERY_MEN_COUNT  int = 2

func main() {
  start := time.Now()
  
  finished := make(chan bool)

  GenerateOrders(ORDERS_COUNT) 
  
  for i := 0; i < ATTENDANTS_COUNT; i++ {
    // 1 second per op
    go Attendant(i + 1)
  }

  for i := 0; i < MACHINES_COUNT; i++ {
    // 2 seconds per op 
    go DoughMachine(i + 1)
  }

  for i := 0; i < BAKERS_COUNT; i++ {
    // 1 second per op 
    go Baker(i + 1)
  }

  for i := 0; i < PACKERS_COUNT; i++ {
    // 500ms per op
    go Packer(i + 1)
  }

  for i := 0; i < DELIVERY_MEN_COUNT; i++ {
    // 2 seconds per op
    go DeliveryMan(i + 1, finished)
  }

  <-finished

  fmt.Println(time.Since(start))
}
