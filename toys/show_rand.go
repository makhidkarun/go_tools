package main

import (
  "fmt"
  mbig  "math/big"
  "math"
  crand "crypto/rand"
)

func main() {
  var all_rolls [7]float64
  var roll_count float64 = 6000
  var mean float64
  var deviation_squares float64
  var off_mean float64
  // var averages float64
  var average_of_deviations float64
  var count_length float64 = float64(len(all_rolls) - 1)
  var i float64

  for i = 0; i < roll_count; i++ {
    roll := RNG(1,6)  
    all_rolls[roll]++
  }
  mean = roll_count/count_length
  for index, value := range all_rolls {
    off_mean = math.Sqrt(float64(mean - value) * (mean - value))
    deviation_squares += off_mean
    if index != 0 {
      fmt.Printf("%d  %.0f %.0f\n", index, value, off_mean)
    }
  }
  average_of_deviations = deviation_squares / count_length
  fmt.Printf("Roll Count: %.0f  ", roll_count)
  fmt.Printf("Mean: %.2f\n", mean)
  //averages = float64(deviation_squares / (len(all_rolls) - 1))
  fmt.Printf("Std Deviation: %.2f \n", math.Sqrt(average_of_deviations))
}

func RNG(min int, max int) int {
  spread := max - min + 1
  spread64 := int64(spread)
  num, _ := crand.Int(crand.Reader, mbig.NewInt(spread64))
  roll := int(num.Int64()) + min 
  return roll
}
