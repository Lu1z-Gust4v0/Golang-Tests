package main

import "errors"

type Menu struct {
  Sizes       map[string]float32
  Flavors     map[string]float32
  Decorations map[string]float32
  Packages    map[string]float32
  Delivery    map[string]float32
}

var menu Menu

func SetUpMenu() {
  menu = Menu {
    Sizes: make(map[string]float32),
    Flavors: make(map[string]float32),
    Decorations: make(map[string]float32),
    Packages: make(map[string]float32),
    Delivery: make(map[string]float32),
  }

  // Price multiplier based on the cake's size
  menu.Sizes["small"]       = 0.5 
  menu.Sizes["normal"]      = 1.0
  menu.Sizes["large"]       = 1.5
  menu.Sizes["extra large"] = 2.0

  menu.Flavors["chocolate"]         = 18.0
  menu.Flavors["strawberry"]        = 18.0
  menu.Flavors["orange"]            = 18.0 
  menu.Flavors["dark forest"]       = 25.0
  menu.Flavors["lemon"]             = 15.0
  menu.Flavors["premium chocolate"] = 30.0 
  menu.Flavors["vanilla"]           = 15.0
  menu.Flavors["carrot"]            = 15.0
  menu.Flavors["blueberry"]         = 20.0
  menu.Flavors["coconut"]           = 15.0
    
  menu.Decorations["chocolate"]         = 5.0
  menu.Decorations["strawberry"]        = 5.0
  menu.Decorations["caramel"]           = 5.0
  menu.Decorations["blueberry"]         = 7.0
  menu.Decorations["vanilla"]           = 5.0
  menu.Decorations["premium chocolate"] = 12.0
  menu.Decorations["dark forest"]       = 10.0
  menu.Decorations["coconut"]           = 5.0
  menu.Decorations["none"]              = 0

  menu.Packages["normal"]   = 2.0
  menu.Packages["birthday"] = 5.0
  menu.Packages["premium"]  = 7.0 
  menu.Packages["golden"]   = 10.0

  menu.Delivery["normal"] = 5.0
  menu.Delivery["fast"]   = 10.0
}

var ORDER_ID_COUNTER uint = 0

func CreateNewOrder(size, flavor, decoration, _package, delivery string) (*Order, error) {
  _, validSize := menu.Sizes[size]
  _, validFlavor := menu.Flavors[flavor]
  _, validDecoration := menu.Decorations[decoration]
  _, validPackage := menu.Packages[_package]
  _, validDelivery := menu.Delivery[delivery]

  if !validSize {
    return nil, errors.New("Invalid size provided")
  }
  
  if !validFlavor {
    return nil, errors.New("Invalid flavor provided")
  }

  if !validDecoration {
    return nil, errors.New("Invalid decoration provided") 
  }

  if !validPackage {
    return nil, errors.New("Invalid package provided")
  }

  if !validDelivery {
    return nil, errors.New("Invalid delivery provided")
  }
  
  ORDER_ID_COUNTER++

  return &Order{
    Id: ORDER_ID_COUNTER,
    Size: size,
    Flavor: flavor, 
    Decoration: decoration,
    Package: _package,
    Delivery: delivery,
  }, nil
} 

func CalculateOrderPrice(order Order) (float32) {
  size := menu.Sizes[order.Size]
  flavor := menu.Flavors[order.Flavor]
  decoration := menu.Decorations[order.Decoration]
  pack := menu.Packages[order.Package]
  delivery := menu.Delivery[order.Delivery]
  
  var price float32 = (size * (flavor + decoration)) + pack + delivery

  return price
}


