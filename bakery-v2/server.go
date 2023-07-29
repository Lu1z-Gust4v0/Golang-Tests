package main

import (
  "github.com/gofiber/fiber/v2"
)

func main() {
  server := fiber.New()
  
  SetUpBakery(10, WorkersConfig{
    AttendantsCount: 2,
    MachinesCount: 4,
    BakersCount: 4,
    PackersCount: 4,
    DeliveryTrucksCount: 2,
  })

  server.Post("/order", MakeNewOrder)

  server.Listen(":8000")
}
