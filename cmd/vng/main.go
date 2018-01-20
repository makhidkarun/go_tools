package main

/* Creates Varg names for the Traveller RPG.
   Traveller is copyright Marc Miller
   This program is provided as is and at no cost,
   under the Far Future Fair Use Policy.
   See http://farfuture.net/FFEFairUsePolicy2008.pdf
*/

import "fmt"       // So we can print purty
import "math/rand" // rand
import "time"      // Seed
import "strings"   // Uppercase first character of name
import "flag"      // Allows --verbose option

var name string = ""
var lastVowel = false
var verbose bool = false

func randInt(min int, max int) int {
	// Remember to give a max that's 1 above actual max
	return min + rand.Intn(max-min)
}

func C() {
	C := randInt(1, 116)
	if verbose {
		fmt.Printf("C is %d.\n", C)
	}

	switch {
	case C <= 5:
		name = name + "d"
	case C <= 10:
		name = name + "dh"
	case C <= 13:
		name = name + "dz"
	case C <= 17:
		name = name + "f"
	case C <= 27:
		name = name + "g"
	case C <= 33:
		name = name + "gh"
	case C <= 35:
		name = name + "gn"
	case C <= 39:
		name = name + "gv"
	case C <= 43:
		name = name + "gz"
	case C <= 53:
		name = name + "k"
	case C <= 56:
		name = name + "kf"
	case C <= 62:
		name = name + "kh"
	case C <= 65:
		name = name + "kn"
	case C <= 68:
		name = name + "ks"
	case C <= 72:
		name = name + "l"
	case C <= 76:
		name = name + "ll"
	case C <= 78:
		name = name + "n"
	case C <= 80:
		name = name + "ng"
	case C <= 85:
		name = name + "r"
	case C <= 89:
		name = name + "rr"
	case C <= 94:
		name = name + "s"
	case C <= 98:
		name = name + "t"
	case C <= 102:
		name = name + "th"
	case C <= 104:
		name = name + "ts"
	case C <= 109:
		name = name + "v"
	case C <= 115:
		name = name + "z"
	}
}

func c() {
	c := randInt(1, 44)
	if verbose {
		fmt.Printf("c is %d.\n", c)
	}

	switch {
	case c <= 1:
		name = name + "dh"
	case c <= 2:
		name = name + "dz"
	case c <= 5:
		name = name + "g"
	case c <= 7:
		name = name + "gh"
	case c <= 8:
		name = name + "ghz"
	case c <= 9:
		name = name + "gz"
	case c <= 11:
		name = name + "k"
	case c <= 13:
		name = name + "kh"
	case c <= 14:
		name = name + "khs"
	case c <= 15:
		name = name + "ks"
	case c <= 17:
		name = name + "l"
	case c <= 18:
		name = name + "ll"
	case c <= 23:
		name = name + "n"
	case c <= 28:
		name = name + "ng"
	case c <= 31:
		name = name + "r"
	case c <= 34:
		name = name + "rr"
	case c <= 35:
		name = name + "rrg"
	case c <= 36:
		name = name + "rrgh"
	case c <= 37:
		name = name + "rs"
	case c <= 38:
		name = name + "rz"
	case c <= 39:
		name = name + "s"
	case c <= 40:
		name = name + "th"
	case c <= 41:
		name = name + "ts"
	case c <= 43:
		name = name + "z"
	}
}

func v() {
	V := randInt(1, 27)
	if verbose {
		fmt.Printf("V is %d.\n", V)
	}

	switch {
	case V <= 5:
		name = name + "a"
	case V <= 9:
		name = name + "ae"
	case V <= 11:
		name = name + "e"
	case V <= 12:
		name = name + "i"
	case V <= 16:
		name = name + "o"
	case V <= 18:
		name = name + "oe"
	case V <= 20:
		name = name + "ou"
	case V <= 23:
		name = name + "u"
	case V <= 26:
		name = name + "ue"
	}
}

func syllable() {
	var z int

	// Implements the note that syllables that end in a vowel are not followed
	// by a syllable that starts with a vowel
	if lastVowel {
		z = randInt(5, 11)
	} else {
		z = randInt(1, 11)
	}

	if verbose {
		fmt.Printf("z is %d.\n", z)
	}

	// C() gives beginning consonants
	// c() gives ending consonants
	// v() gives vowels
	// This switch statement follows the math but not the form
	// of the chart under "Language and Naming". While it gives
	// The same result it switches VC and CV to allow for
	// lastVowel to be easier to use.

	switch {
	case z <= 1:
		v()
		lastVowel = true
	case z <= 4:
		v()
		c()
		lastVowel = false
	case z <= 7:
		C()
		v()
		lastVowel = true
	case z <= 10:
		C()
		v()
		c()
		lastVowel = false
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	flag.BoolVar(&verbose, "verbose", false, "Show results")
	flag.Parse()

	// Vargr names are rarely longer than 6 syllables
	nameLength := randInt(1, 7)
	if verbose {
		fmt.Printf("Name length is %d.\n", nameLength)
	}

	for i := 0; i < nameLength; i++ {
		syllable()
	}

	fmt.Println(strings.Title(name))
}
