// Functions for Traveller and Cepheus Engine
// All copyrights to them that own's them.

package go_tools

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	crand "crypto/rand"
	mbig "math/big"

	_ "github.com/mattn/go-sqlite3"
)

func Age(terms int) int {
	return 18 + (terms * 4) + RNG(0, 3)
}

func ArrayFromFile(f *os.File) []string {
	items := make([]string, 0)
	input := bufio.NewScanner(f)
	for input.Scan() {
		s := input.Text()
		if len(s) > 2 {
			items = append(items, s)
		}
	}
	return items
}

func RNG(min int, max int) int {
	spread := max - min + 1
	spread64 := int64(spread)
	num, _ := crand.Int(crand.Reader, mbig.NewInt(spread64))
	roll := int(num.Int64()) + min
	return roll
}

func OneD6() int {
	return RNG(1, 6)
}

func TwoD6() int {
	return OneD6() + OneD6()
}

func FormatUPP(upp [6]int) string {
	var newUPP string
	for _, val := range upp {
		newUPP += fmt.Sprintf("%X", val)
	}
	return newUPP
}

func Gender() string {
	if RNG(1, 6)%2 == 0 {
		return "F"
	} else {
		return "M"
	}
}

func GetName(gender string, db_name string) string {
	// Note that the names.db file must be where the command is run
	// from.

	var lname string
	var fname string
	var fresult *sql.Rows

	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		fmt.Println(err)
	}

	// First Name
	if gender == "M" {
		fresult, err = db.Query("SELECT name FROM humaniti_male_first ORDER BY RANDOM() LIMIT 1")
	} else {
		fresult, err = db.Query("SELECT name FROM humaniti_female_first ORDER BY RANDOM() LIMIT 1")
	}
	if err != nil {
		fmt.Println(err)
	}
	for fresult.Next() {
		err = fresult.Scan(&fname)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Last Name
	result, err := db.Query("SELECT name FROM humaniti_last ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		fmt.Println(err)
	}
	for result.Next() {
		err = result.Scan(&lname)
		if err != nil {
			fmt.Println(err)
		}
	}

	db.Close()
	name := fname + " " + lname
	return name
}

func modifyUpp(upp [6]int, stat int, mod int) [6]int {
	// Requires the UPP [6]int, stat index [0-5], and modifier
	if stat < 0 || stat > 5 {
		return upp
	} else {
		upp[stat] += mod
		if upp[stat] < 0 {
			upp[stat] = 0
		} else if upp[stat] > 15 {
			upp[stat] = 15
		}
	}
	return upp
}

func NumTerms() int {
	return RNG(1, 7)
}

func RollUPP() [6]int {
	var upp [6]int
	for i := 0; i < 6; i++ {
		upp[i] = TwoD6()
	}
	return upp
}

func StringInArray(val string, array []string) bool {
	for _, value := range array {
		if value == val {
			return true
		}
	}
	return false
}

func RandomStringFromArray(array []string) string {
	return array[RNG(0, len(array)-1)]
}
