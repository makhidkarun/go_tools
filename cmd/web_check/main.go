// web_check
// Used to test URLs to ensure they have relevant code.
// Starting from the example on page 93 of Caleb Doxsey's
// "Introducing Go".

// TODO:
// 1.  Make it work. DONE
// 2.  Grab response codes. DONE
// 3.  Check for key words in response.Body. DONE
// 4.  Take a file of URLs. DONE
// 4.1  Error checking.
// 4.2  Check for and add http:// if missing. Or work around.
// 5.  Add key words to data file. (?csv)
// 6.  Set up base round of tests.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type HomePageSize struct {
	URL      string
	Size     int
	Response int
	Body     []byte
}

var file = flag.String("f", "-", "File to read")

func has_string(body []byte, term string) bool {
	body_string := string(body[:len(body)])
	return strings.Contains(body_string, term)
}

func lines(file *os.File) (lines []string) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		my_string := input.Text()
		if len(my_string) > 2 {
			lines = append(lines, my_string)
		}
	}
	return
}
func main() {
	flag.Parse()
	data, err := os.Open(*file)
	if err != nil {
		fmt.Fprint(os.Stderr, "Open failed: %v\n", err)
		os.Exit(1)
	}
	defer data.Close()

	urls := lines(data)

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			resp := res.StatusCode

			results <- HomePageSize{
				URL:      url,
				Size:     len(bs),
				Response: resp,
				Body:     bs,
			}
		}(url)
	}

	for range urls {
		result := <-results
		fmt.Println("For ", result.URL, "the response was: ", result.Response)
		fmt.Println(has_string(result.Body, "doufdofudoifud")) // Should give false
		fmt.Println(has_string(result.Body, "google"))         // Should usually give true
	}

}
