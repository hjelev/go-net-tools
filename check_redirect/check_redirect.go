package main

import (
    "errors"
    "fmt"
    "net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		domain := "https://" + os.Args[1]
		req, err := http.NewRequest("GET", domain, nil)
		if err != nil {
			panic(err)
		}
		client := new(http.Client)
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return errors.New("Redirect")
		}

		response, err := client.Do(req)
		if err == nil {
			fmt.Println(response.StatusCode, domain)

		} else {
			
			if strings.Contains(err.Error(), "Redirect") {
				fmt.Println(response.StatusCode, response.Header.Get("Location"))
			} else { 
				fmt.Println(err) 
			}
		}

	} else {
		fmt.Printf("Usage: check_redirect domain.tld")
	}
}