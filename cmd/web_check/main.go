// web_check.go
// Used to test URLs to ensure they have relevant code.
// Starting from the example on page 93 of Caleb Doxsey's 
// "Introducing Go".

// TODO:
// 1.  Make it work. DONE
// 2.  Grab response codes. DONE
// 3.  Check for key words in response.Body.
// 4.  Take a file of URLs and key words
// 5.  Set up base round of tests.

package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
)

type HomePageSize struct {
  URL  			string
  Size 			int
  Response  int
	Body			[]byte
}

func has_string(body []byte, term string) bool {
	body_string := string(body[:len(body)])
	return strings.Contains(body_string, term)
}

func main() {
  urls := []string{
    "http://www.apple.com",
    "http://www.amazon.com",
    "http://www.google.com",
    "http://microsoft.com",
  }

  results := make(chan HomePageSize)

  for _,url := range urls {
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
        URL:  url,
        Size: len(bs),
        Response: resp,
				Body:			bs,
      }
    }(url)
  }

  for range urls {
    result := <- results
		fmt.Println("For ", result.URL, "the response was: ", result.Response)
		fmt.Println(has_string(result.Body, "doufdofudoifud"))  // Should give false
		fmt.Println(has_string(result.Body, "google")) 				  // Should usually give true
  }
}
