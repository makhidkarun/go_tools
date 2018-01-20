package main

import (
  "fmt"
  "os"
  binary "encoding/binary"
  crand "crypto/rand"
  mrand "math/rand"
)

func Seed() mrand.Source {
  var seed int64
  err := binary.Read(crand.Reader, binary.BigEndian, &seed)
  if err != nil {
   fmt.Println(err)
   }
 return mrand.NewSource(seed)
}

func main(){
  seed := Seed()
  rng := mrand.New(seed)
  for i := 0; i < 1000; i++ {
    roll := rng.Intn(6)
    if roll == 0 {
      fmt.Println("Have a 0!")
      os.Exit(0)
    }
  }
}


