package main

import (
    "net/http"
    "os"
    "fmt"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a URL as an argument.")
        fmt.Println("Example: cookies http://example.com")
        return
    }
    url := os.Args[1]
    if !strings.HasPrefix(url, "http") {
        url = "https://" + url
    }
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    for _, cookie := range resp.Cookies() {
        fmt.Printf("Name: %s, Value: %s\n", cookie.Name, cookie.Value)
    }
}