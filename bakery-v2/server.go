package main

import (
  "github.com/gofiber/fiber/v2"
)

func main() {
  server := fiber.New()
  
  SetUpBakery(20, WorkersConfig{
    Limiters: 2,
    AttendantsCount: 15,
    MachinesCount: 30,
    BakersCount: 30,
    PackersCount: 30,
    DeliveryTrucksCount: 15,
  })

  server.Post("/order", MakeNewOrder)
  server.Get("/balance", CheckBakeryBalance)

  server.Listen(":8000")
}
