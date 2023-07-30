package main

import (
	"github.com/gofiber/fiber/v2"
) 

type NewOrderRequestBody struct {
  Size        string `json:"size"       validation:"required"` 
  Flavor      string `json:"flavor"     validation:"required"`
  Decoration  string `json:"decoration" validation:"required"` 
  Package     string `json:"package"    validation:"required"` 
  Delivery    string `json:"delivery"   validation:"required"`
}

func MakeNewOrder(context *fiber.Ctx) error {
  newOrder := new(NewOrderRequestBody)

  if error := context.BodyParser(newOrder); error != nil {
    return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message": "Invalid Order",
      "error": error,
    }) 
  } 

  order, error := CreateNewOrder( 
    newOrder.Size, 
    newOrder.Flavor, 
    newOrder.Decoration, 
    newOrder.Package, 
    newOrder.Delivery,
  ) 
  
  if error != nil {
    return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
      "message": "Invalid Order",
      "order": nil,
      "error": error,
    })
  }
  
  if len(bakeryChannels.PendingOrders) == cap(bakeryChannels.PendingOrders) {
    LimiterQueue <- *order
    
    return context.Status(fiber.StatusOK).JSON(fiber.Map{
      "message": "Order completed sucessfully",
      "order": *order,
      "error": nil,
    })
  }  

  bakeryChannels.PendingOrders <- *order

  return context.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "Order completed sucessfully",
    "order": *order,
    "error": nil,
  })
}

func CheckBakeryBalance(context *fiber.Ctx) error {
  balance := bakeryState.Balance 
  
  if balance == 0 {
    return context.Status(fiber.StatusOK).JSON(fiber.Map{
      "message": "Unfortunately we ran out of money",
      "balance": balance,
      "error": nil,
    })
  }

  return context.Status(fiber.StatusOK).JSON(fiber.Map{
    "message": "Bakery's balance fetched sucessfully",
    "balance": balance,
    "error": nil,
  })
}
