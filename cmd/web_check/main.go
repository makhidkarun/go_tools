// web_check.go
// Used to test URLs to ensure they have relevant code.
// Starting from the example on page 93 of Caleb Doxsey's 
// "Introducing Go".

// TODO:
// 1.  Make it work.
// 2.  Grab response codes.
// 3.  Check for key words in response.Body.
// 4.  Take a file of URLs and key words

package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

type HomePageSize struct {
  URL  			string
  Size 			int
  Response  int
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
      }
    }(url)
  }

  //var biggest HomePageSize

  for range urls {
    result := <- results
		fmt.Println("For ", result.URL, "the response was: ", result.Response)
    //result := <- results
    //if result.Size > biggest.Size {
    //  biggest = result
    //}
  }

  //fmt.Println("The biggest home page: ", biggest.URL)
  //fmt.Println("The response was: ", biggest.Response)
}
