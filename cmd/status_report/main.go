// Used to create a set of rolls from a list of strings in a file.
// For example, roll to see how well each NPC did that quarter.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/makhidkarun/go_tools"
)

var file = flag.String("f", "-", "File to read.")
var count = flag.Int("c", 1, "Count of rolls per line.")
var sep = flag.String("s", " ", "Field separator.")

func main() {
	flag.Parse()
	data, err := os.Open(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open failed: %v\n", err)
		os.Exit(1)
	}
	printLines(data, *count, *sep)
	data.Close()
}

func printLines(f *os.File, count int, sep string) {
	input := bufio.NewScanner(f)
	var rolls []string
	for input.Scan() {
		my_string := input.Text()
		if len(my_string) > 2 {
			for i := 0; i < count; i++ {
				rolls = append(rolls, strconv.Itoa(go_tools.TwoD6()))
			}
			fmt.Printf("%s%s%s\n", my_string, sep, strings.Join(rolls, sep))
			rolls = nil
		}
	}
}
