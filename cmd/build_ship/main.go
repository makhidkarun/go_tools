// build_ship.go

package main

import (
  "fmt"
  //"flag"
)

type Ship struct {
  name        string
  hull_size   int
  drive_size  int
}

func main() {
  miss_rosa := Ship{name: "Miss Rosa", hull_size: 200, drive_size: 35}
  fmt.Printf("The %s is a %d dTon ship with %d of drives.\n", 
    miss_rosa.name, miss_rosa.hull_size, miss_rosa.drive_size)

}
