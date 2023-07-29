package main

import (
  "github.com/gofiber/fiber/v2"
) 

func MakeNewOrder(context *fiber.Ctx) error {
  order := new(Order)

  if error := context.BodyParser(order); error != nil {
    return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message": "Invalid Order",
      "error": error,
    }) 
  } 

  bakeryState.PendingOrders += 1
  bakeryChannels.PendingOrders <- Order{
    Id: order.Id,
    Flavor: order.Flavor,
    Decoration: order.Decoration,
    Package: order.Package,
    Delivery: order.Delivery,
    Total: order.Total,
  } 

  return context.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "Order completed sucessfully",
    "error": nil,
  })
}
