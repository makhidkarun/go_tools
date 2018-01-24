
package main

import (
  "flag"
  "fmt"
 
  "github.com/makhidkarun/go_tools"
)

func main() {
  var num_rolls = flag.Int("n", 1, "Number of rolls")
  flag.Parse()

  for i := 0; i < *num_rolls; i++ {
    fmt.Printf("%d%d\n", go_tools.OneD6(), go_tools.OneD6())
  }
}


