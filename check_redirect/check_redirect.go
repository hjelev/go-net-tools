package main

import (
    "errors"
    "fmt"
    "net/http"
	"os"
	"strings"
)

func check_redirect(url string)(string, int){

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}

	response, err := client.Do(req)
	if err == nil {
		fmt.Println(response.StatusCode, url)

	} else {
		
		if strings.Contains(err.Error(), "Redirect") {
			fmt.Println(response.StatusCode, url)
		} else { 
			fmt.Println(err) 
		}
	}
	return response.Header.Get("Location"), response.StatusCode
}

func main() {
	if len(os.Args) > 1 {
		var domain string
		if strings.HasPrefix(os.Args[1], "https://") || strings.HasPrefix(os.Args[1], "http://") {
			domain = os.Args[1]
		} else {
			domain = "https://" + os.Args[1]
		}		
		for {
			resp_url, resp_code := check_redirect(domain)
			domain = resp_url
	
			if resp_code < 300 || resp_code > 400 {
				break
			}
		}
	} else {
		fmt.Printf("Usage: check_redirect domain.tld \n")
	}
}
