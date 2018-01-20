package go_tools

import (
  mrand "math/rand"
  "testing"
)

var seed = Seed()
var rng  = mrand.New(seed)

func TestOneD6(t *testing.T){
  roll := OneD6(rng)
  if roll < 1 || roll > 6 {
    t.Error(`OneD6 failed`)
  }
}

func TestTwoD6(t *testing.T){
  roll := TwoD6(rng)
  if roll < 2 || roll > 12 {
    t.Error(`TwoD6 failed`)
  }
}

func TestAge(t *testing.T){
  age := Age(2, rng)
  if age < 26 || age > 29 {
    t.Error(`Age failed`)
  }
}

func TestFormatUPP(t *testing.T){
  upp := [6]int{ 7, 7, 15, 9, 12, 12}
  newUPP := FormatUPP(upp)
  if newUPP != "77F9CC" {
    t.Error(`FormatUPP failed.`)
  }
}

func TestStringInArray(t *testing.T){
  var genders []string = []string{"F", "M"}
  option1 := "F"
  option2 := "R"
  if !StringInArray(option1, genders){
    t.Error(`Missing an F.`)
  }
  if StringInArray(option2, genders){
    t.Error(`Never met R.`)
  }
}

func TestRandomStringFromArray(t *testing.T){
  var genders []string = []string{"F", "M"}
  gender := RandomStringFromArray(genders)
  if !StringInArray(gender, genders){
    t.Error(`Bad gender output.`)
  }
  var ranks []string = []string{"PVT", "SGT", "LT", "CPT", "MAJ"}
  rank := RandomStringFromArray(ranks)
  if !StringInArray(rank, ranks){
    t.Error(`Bad rank output.`)
  }
}
