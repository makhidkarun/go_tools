// To provide two names.

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
	//var db_name = flag.String("d", "data/names.db", "Database")

	//flag.Parse()

	// This will need a lot of testing.
	_, err := os.Stat(*db_name)
	if err != nil {
		fmt.Println(err)
	}
	options["db_name"] = *db_name

	//var character go_tools.Person = go_tools.MakePerson(options)

	//show_upp := go_tools.FormatUPP(character.UPP)

	fmt.Printf("%s %s [%s] (age %d)\n",
		character.Name, show_upp, character.Gender, character.Age)
}
