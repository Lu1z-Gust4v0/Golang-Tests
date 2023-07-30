package main

import (
  "time"
  "fmt"
)


type WorkersConfig struct {
  Limiters            int `json:"limiters"`
  AttendantsCount     int `json:"attendants"`
	MachinesCount       int `json:"machines"`
	BakersCount         int `json:"bakers"`
	PackersCount        int `json:"packers"`
	DeliveryTrucksCount int `json:"trucks"`
}

func TimeNow() string {
  now := time.Now()

  return fmt.Sprintf("%02d/%02d/%d - %02d:%02d:%02d",
    now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second(),
  )
}

func SetUpWorkers(config WorkersConfig) {
  for i := 0; i < config.Limiters; i++ {
    go Limiter(i + 1)
  }

	for i := 0; i < config.AttendantsCount; i++ {
		go Attendant(i + 1)
	}

	for i := 0; i < config.MachinesCount; i++ {
		go DoughMachine(i + 1)
	}

	for i := 0; i < config.BakersCount; i++ {
		go Baker(i + 1)
	}

	for i := 0; i < config.PackersCount; i++ {
		go Packer(i + 1)
	}

	for i := 0; i < config.DeliveryTrucksCount; i++ {
		go DeliveryTruck(i + 1)
	}
}

var LimiterQueue = make(chan Order, 100)

func Limiter(index int) {
  for request := range LimiterQueue {
    fmt.Printf("[Limiter %v] \tRequest %v is being limited \t[%v]\n", 
      index, request.Id, TimeNow(),
    )
    time.Sleep(200 * time.Millisecond) 
    bakeryChannels.PendingOrders <- request
  } 
} 

func Attendant(index int) {
	for order := range bakeryChannels.PendingOrders {
		var cake = Cake{
			Order: order,
			State: "ordered",
		}
    
    fmt.Printf("[Attndt %v]: \t[Order %v] Ordering \t%v %v cake \t[%v]\n", 
      index, order.Id, order.Size, order.Flavor, TimeNow(),
    )

		time.Sleep(500 * time.Millisecond)

    bakeryChannels.CakesToBake <- cake
	
    bakeryState.mutex.Lock()
    bakeryState.PendingOrders -= 1
    bakeryState.CakesToBake += 1
    bakeryState.mutex.Unlock()
  }
}

func DoughMachine(index int) {
  for cake := range bakeryChannels.CakesToBake {
    cake.State = "baked"
   
    fmt.Printf("[Machine %v]: \t[Order %v] Baking \t%v %v cake \t[%v]\n",
      index, cake.Order.Id, cake.Order.Size,cake.Order.Flavor, TimeNow(),
    )
    
    time.Sleep(2 * time.Second)

    bakeryChannels.CakesToDecorate <- cake 
    
    bakeryState.mutex.Lock()
    bakeryState.CakesToBake -= 1
    bakeryState.CakesToDecorate += 1
    bakeryState.mutex.Unlock()
  }
}

func Baker(index int) {
  for cake := range bakeryChannels.CakesToDecorate {
    cake.State = "decorated"
   
    fmt.Printf("[Baker %v]: \t[Order %v] Decorating \t%v %v cake \t[%v]\n",
      index, cake.Order.Id, cake.Order.Size, cake.Order.Flavor, TimeNow(),
    )

    time.Sleep(time.Second)

    bakeryChannels.CakesToPack <- cake 
    
    bakeryState.mutex.Lock()
    bakeryState.CakesToDecorate -= 1
    bakeryState.CakesToPack += 1
    bakeryState.mutex.Unlock()
  }
}

func Packer(index int) {
  for cake := range bakeryChannels.CakesToPack {
    cake.State = "packed"
   
    fmt.Printf("[Packer %v]: \t[Order %v] Packing \t%v %v cake \t[%v]\n",
      index, cake.Order.Id, cake.Order.Size, cake.Order.Flavor, TimeNow(),
    )

    time.Sleep(500 * time.Millisecond)

    bakeryChannels.CakesToDeliver <- cake 
    
    bakeryState.mutex.Lock()
    bakeryState.CakesToPack -= 1
    bakeryState.CakesToDeliver += 1
    bakeryState.mutex.Unlock()
  }
}

func DeliveryTruck(index int) {
  for cake := range bakeryChannels.CakesToDeliver {
    cake.State = "delivered"
   
    fmt.Printf("[Truck %v]: \t[Order %v] Delivering \t%v %v cake \t[%v]\n",
      index, cake.Order.Id, cake.Order.Size, cake.Order.Flavor, TimeNow(),
    )
    
    if cake.Order.Delivery == "fast" {
      time.Sleep(time.Second)
    } else {
      time.Sleep(2 * time.Second)
    } 

    fmt.Printf("[Truck %v]: \t[Order %v] Delivered \t%v %v cake \t[%v]\n",
      index, cake.Order.Id, cake.Order.Size, cake.Order.Flavor, TimeNow(),
    )

    bakeryState.mutex.Lock()
    bakeryState.CakesToDeliver -= 1
    bakeryState.Balance += CalculateOrderPrice(cake.Order)
    bakeryState.mutex.Unlock()
  }
}
