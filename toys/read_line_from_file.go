package main

import (
  "bufio"
  "flag"
  "fmt"
  "os"

  "github.com/makhidkarun/go_tools"
)

var file = flag.String("f", "-", "File to read.")

func main() {
  flag.Parse()
  data, err := os.Open(*file)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Open failed: %v\n", err)
    os.Exit(1)
  }
  randomLine(data)
  data.Close()
}

func randomLine(f *os.File) {
  words := make([]string, 0)
  input := bufio.NewScanner(f)
  for input.Scan() {
    my_string := input.Text()
    if len(my_string) > 2 {
      words = append(words, my_string)
    }
  }
  index := go_tools.RNG(1, len(words) - 1)
  fmt.Println(words[index])
}
