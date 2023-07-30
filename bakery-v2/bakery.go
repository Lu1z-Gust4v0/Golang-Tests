package main

import (
  "sync"
)

type BakeryState struct {
  mutex             sync.Mutex
  Balance           float32
  PendingOrders     int
  CakesToBake       int 
  CakesToDecorate   int
  CakesToPack       int
  CakesToDeliver    int
}

type BakeryChannels struct {
  PendingOrders     chan Order 
  CakesToBake       chan Cake 
  CakesToDecorate   chan Cake
  CakesToPack       chan Cake
  CakesToDeliver    chan Cake
}

type Order struct {
  Id          uint    `json:"id"          validate:"required"` 
  Size        string  `json:"size"        validate:"required"`
  Flavor      string  `json:"flavor"      validate:"required"` 
  Decoration  string  `json:"decoration"  validate:"required"` 
  Package     string  `json:"package"     validate:"required"`
  Delivery    string  `json:"delivery"    validate:"required"` 
}

type Cake struct {
  Order   Order 
  State   string
}

var bakeryState BakeryState 
var bakeryChannels BakeryChannels

func SetUpBakery(capacity int, config WorkersConfig) {
  bakeryState = BakeryState {
    Balance: 0, 
    PendingOrders: 0,
    CakesToBake: 0,
    CakesToDecorate: 0,
    CakesToPack: 0,
    CakesToDeliver: 0,
  } 

  bakeryChannels = BakeryChannels {
    PendingOrders: make(chan Order, capacity),
    CakesToBake: make(chan Cake, capacity),
    CakesToDecorate: make(chan Cake, capacity),
    CakesToPack: make(chan Cake, capacity),
    CakesToDeliver: make(chan Cake, capacity),
  }
  
  SetUpMenu()
  SetUpWorkers(config)
}
