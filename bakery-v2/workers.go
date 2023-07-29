package main

import (
  "time"
  "fmt"
)


type WorkersConfig struct {
	AttendantsCount     int `json:"attendants"`
	MachinesCount       int `json:"machines"`
	BakersCount         int `json:"bakers"`
	PackersCount        int `json:"packers"`
	DeliveryTrucksCount int `json:"trucks"`
}

func TimeNow() string {
  now := time.Now()

  return fmt.Sprintf("%d/%d/%d - %d:%d:%d",
    now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second(),
  )
}

func SetUpWorkers(config WorkersConfig) {
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

func Attendant(index int) {
	for order := range bakeryChannels.PendingOrders {
		var cake = Cake{
			Order: order,
			State: "ordered",
		}
    
    fmt.Printf("Attendant [%v]: [%v] Ordering a %v cake decorated with %v at %v\n", 
      index, order.Id, order.Flavor, order.Decoration, TimeNow(),
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
   
    fmt.Printf("Machine [%v]: [%v] Baking a %v cake at %v\n",
      index, cake.Order.Id, cake.Order.Flavor, TimeNow(),
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
   
    fmt.Printf("Baker [%v]: [%v] Decorating a %v cake with %v at %v\n",
      index, cake.Order.Id, cake.Order.Flavor, cake.Order.Decoration, TimeNow(),
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
   
    fmt.Printf("Packer [%v]: [%v] Packing a %v cake at %v\n",
      index, cake.Order.Id, cake.Order.Flavor, TimeNow(),
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
   
    fmt.Printf("Truck [%v]: [%v] delivering a %v cake at %v\n",
      index, cake.Order.Id, cake.Order.Flavor, TimeNow(),
    )
    
    if cake.Order.Delivery == "fast" {
      time.Sleep(time.Second)
    } else {
      time.Sleep(2 * time.Second)
    } 

    fmt.Printf("Truck [%v]: Cake [%v] delivered at %v\n",
      index, cake.Order.Id, TimeNow(),
    )

    bakeryState.mutex.Lock()
    bakeryState.CakesToDeliver -= 1
    bakeryState.Balance += cake.Order.Total
    bakeryState.mutex.Unlock()
  }
}
