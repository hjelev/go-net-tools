package main

import (
    "errors"
    "fmt"
    "net/http"
	"os"
	"strings"
)

func handlePanic(url string) {
	// detect if panic occurs or not
	a := recover()
	if a != nil {
	  fmt.Println(url, " is not a valid url!")
	}
}

func check_redirect(url string)(string, int){
	defer handlePanic(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		os.Exit(1)
	}

	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}

	response, err := client.Do(req)

	fmt.Println(response.StatusCode, url)

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
		fmt.Println(`Usage: check_redirect [URL]
Display status code returned from url and follow redirects.
Automatically appends https:// to url if not present.`)
	}
}