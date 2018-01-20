// To create a basic 2d6 OGL character

// With acknowledgement of Marc Miller of FFE, 
// Jason "Flynn" Kemp of Cepheus Engine,
// Rob Pike and the Google Go team,
// and Freenode#go-nuts.

package go_tools

import (
  "strconv"
  "strings"
)

type Person struct {
  Name    string
  UPP     [6]int
  Gender  string
  PSR     int
  Age     int
  Terms   int
  Career  string
}

func MakePerson(options map[string]string) Person {
  terms, _  := strconv.Atoi(options["terms"])
  gender    := options["gender"]
  var character Person 

  // Need to figure out pre-adult characters
  if terms <= 0 {
    character.Terms = NumTerms()
  } else {
    character.Terms = terms
  }

  var genders []string = []string{"F", "M"}
  input_gender := strings.ToUpper(gender)
  if !StringInArray(input_gender, genders) {
    character.Gender = Gender()
  } else {
    character.Gender = input_gender
  }

  character.Name    = GetName(character.Gender)
  character.UPP     = RollUPP()
  character.Age     = Age(character.Terms)

  return character
}
