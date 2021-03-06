// To create a basic 2d6 OGL character

// With acknowledgement of Marc Miller of FFE,
// Jason "Flynn" Kemp of Cepheus Engine
//  Rob Pike and the Google Go team,
//  and Freenode#go-nuts.

package main

import (
	"flag"
	"fmt"
	"github.com/makhidkarun/go_tools"
	"os"
	"strconv"
	"strings"
)

func main() {

	var options = make(map[string]string)
	var terms = flag.Int("t", 0, "Number of terms")
	var gender = flag.String("g", "", "Gender")
	var db_name = flag.String("d", "data/names.db", "Database")

	flag.Parse()

	// This will need a lot of testing.
	_, err := os.Stat(*db_name)
	if err != nil {
		fmt.Println(err)
	}
	options["db_name"] = *db_name

	// Need to figure out pre-adult characters
	if *terms > 0 && *terms <= 10 {
		options["terms"] = strconv.Itoa(*terms)
	}

	// Needs some input checking.
	var genders []string = []string{"F", "M"}
	input_gender := strings.ToUpper(*gender)
	if go_tools.StringInArray(input_gender, genders) {
		options["gender"] = input_gender
	}

	var character go_tools.Person = go_tools.MakePerson(options)

	show_upp := go_tools.FormatUPP(character.UPP)

	fmt.Printf("%s %s [%s] (age %d)\n",
		character.Name, show_upp, character.Gender, character.Age)
}
