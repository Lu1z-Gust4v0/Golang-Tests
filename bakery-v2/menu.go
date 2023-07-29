package main

type Menu struct {
  flavors     map[string]int 
  decorations map[string]int
  packages    map[string]int
  delivery    map[string]int
}

var menu Menu

func InitMenu() {
  menu = Menu {
    flavors: make(map[string]int),
    decorations: make(map[string]int),
    packages: make(map[string]int),
    delivery: make(map[string]int),
  }

  menu.flavors["chocolate"] = 18 
  menu.flavors["strawberry"] = 18
  menu.flavors["orange"] = 18 
  menu.flavors["dark_forest"] = 25
  menu.flavors["lemon"] = 15
  menu.flavors["premium_chocolate"] = 30 
  menu.flavors["vanilla"] = 15
  
  menu.decorations["chocolate"] = 5
  menu.decorations["strawberry"] = 5
  menu.decorations["vanilla"] = 5
  menu.decorations["premium_chocolate"] = 10
  menu.decorations["dark_forest"] = 8
  menu.decorations["none"] = 0

  menu.packages["normal"] = 2
  menu.packages["birthday"] = 5
  menu.packages["premium"] = 7

  menu.delivery["normal"] = 5
  menu.delivery["fast"] = 10
}
