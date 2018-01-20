package main

import (
  "fmt"
  "strings"
)

func StringInArray(val string, array []string) bool {
  for _, value := range array {
    if value == val {
      return true
    }
  }
  return false
} 

func main() {
  var name string = "Fred"
  name_print := strings.ToUpper(name)
  fmt.Println(name_print)

  var genders []string = []string{"F", "M"}
  if StringInArray("F", genders) {
    fmt.Println("Found a Female!")
  }
  if StringInArray("R", genders) {
    fmt.Println("Need to fix the function.")
  } 
}

