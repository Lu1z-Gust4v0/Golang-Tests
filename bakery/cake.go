package main

import (
	"fmt"
	"time"
)

type Bakery struct {
  PendingOrders     int
  OrdersToConfirm   int
  CakesToBake       int
  CakesToDecorate   int
  CakesToPack       int
  CakesToDeliver    int
  channels BakeryChannels
}

type BakeryChannels struct {
  Orders          chan int
  ConfirmedOrders chan int 
  BakedCakes      chan int
  DecoratedCakes  chan int
  PackedCakes     chan int
  DelivedCakes    chan int
}

var bakery Bakery

func order(cake int) {
  fmt.Printf("Receiving order - %v\n", cake)
  time.Sleep(time.Second)
  fmt.Printf("Order sent - %v\n", cake)
  bakery.channels.ConfirmedOrders <- cake
  bakery.OrdersToConfirm -= 1
  
  if bakery.OrdersToConfirm == 0 {
    close(bakery.channels.ConfirmedOrders)
  } 
}

func bake(cake int) {
  fmt.Printf("Baking cake - %v\n", cake)
  time.Sleep(2 * time.Second)
  fmt.Printf("Cake baked - %v\n", cake)
  bakery.channels.BakedCakes <- cake
  bakery.CakesToBake -= 1

  if bakery.CakesToBake == 0 {
    close(bakery.channels.BakedCakes)
  }
}

func decorate(cake int) {
  fmt.Printf("Decorating cake - %v\n", cake)
  time.Sleep(time.Second)
  fmt.Printf("Cake decorated - %v\n", cake)
  bakery.channels.DecoratedCakes <- cake
  bakery.CakesToDecorate -= 1

  if bakery.CakesToDecorate == 0 {
    close(bakery.channels.DecoratedCakes)
  }
}

func pack(cake int) {
  fmt.Printf("Packing cake - %v\n", cake)
  time.Sleep(500 * time.Millisecond)
  fmt.Printf("Cake packed - %v\n", cake)
  bakery.channels.PackedCakes <- cake
  bakery.CakesToPack -= 1

  if bakery.CakesToPack == 0 {
    close(bakery.channels.PackedCakes)
  }
}

func deliver(cake int) {
  fmt.Printf("Delivering cake - %v\n", cake)
  time.Sleep(2 * time.Second)
  fmt.Printf("Cake delivered - %v\n", cake)
  bakery.channels.DelivedCakes <- cake
  bakery.CakesToDeliver -= 1
  
  if bakery.CakesToDeliver == 0 {
    close(bakery.channels.DelivedCakes)
  }
}

func GenerateOrders(count int) {
  bakery.PendingOrders = count
  bakery.OrdersToConfirm = count
  bakery.CakesToBake = count
  bakery.CakesToDecorate = count 
  bakery.CakesToPack = count
  bakery.CakesToDeliver = count

  bakery.channels = BakeryChannels{
    Orders: make(chan int, count),
    ConfirmedOrders: make(chan int, count),
    BakedCakes: make(chan int, count),
    DecoratedCakes: make(chan int, count),
    PackedCakes: make(chan int, count),
    DelivedCakes: make(chan int, count),
  }
  
  defer close(bakery.channels.Orders)
  
  for i := 0; i < count; i++ {
    bakery.channels.Orders <- i + 1
  }
}

func Attendant(number int) {
  for newOrder := range bakery.channels.Orders {
    fmt.Printf("Attendant %v is recieving an order\n", number)
    order(newOrder)
  }
}

func DoughMachine(number int) {
  for order := range bakery.channels.ConfirmedOrders {
    fmt.Printf("Dough machine %v is active\n", number)
    bake(order)
  }
}

func Baker(number int) {
  for cake := range bakery.channels.BakedCakes {
    fmt.Printf("Baker %v is decorating a cake\n", number)
    decorate(cake)
  }
}

func Packer(number int) {
  for cake := range bakery.channels.DecoratedCakes {
    fmt.Printf("Packer %v is packing a cake\n", number)   
    pack(cake)
  }
}

func DeliveryMan(number int, finished chan bool) {
  defer func () {
    finished <- true
  }() 
  
  for packed := range bakery.channels.PackedCakes {
    fmt.Printf("Delivery man %v is delivery a cake\n", number)  
    deliver(packed)
  }
}
