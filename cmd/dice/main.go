// dice.go
// General dice roller
// 
//  TODO:
//  1. Parse line into array of strings, based on ',' split.
//  2. Parse each line into numDice, numSides, modifier (2d6+3)
//  3. Write test for random generation.
//  4. Generate random number.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Insert test for lack of os.Args[1]
	rolls 	:= strings.Split(os.Args[1], ",")
	for _, roll := range rolls {
		fmt.Println(roll)
	}
}
