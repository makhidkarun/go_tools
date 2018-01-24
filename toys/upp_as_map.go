package main

import (
  "fmt"

  "github.com/makhidkarun/go_tools"
)

func main() {
  upp := make(map[string]int)
  upp["str"]  = go_tools.TwoD6()
  upp["dex"]  = go_tools.TwoD6()
  upp["end"]  = go_tools.TwoD6()
  upp["int"]  = go_tools.TwoD6()
  upp["edu"]  = go_tools.TwoD6()
  upp["soc"]  = go_tools.TwoD6()
  fmt.Println(upp["str"])
  fmt.Println(upp["dex"])
  fmt.Println(upp["end"])
  fmt.Println(upp["int"])
  fmt.Println(upp["edu"])
  fmt.Println(upp["soc"])
}


