package main

import (
	"crypto/tls"
	"fmt"
	"time"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		domain := os.Args[1]
		conn, err := tls.Dial("tcp", domain + ":443", nil)
		if err != nil {
			fmt.Println("Server doesn't support SSL certificate err: " + err.Error())
		} else {

			err = conn.VerifyHostname(domain)
			if err != nil {
				panic("Hostname doesn't match with certificate: " + err.Error())
			}
			
			expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
			fmt.Printf("Issuer: %s\nExpiry: %v\n", conn.ConnectionState().PeerCertificates[0].Issuer, expiry.Format(time.RFC850))
		}
	} else {
		fmt.Println("Usage: check_ssl domain.tld")
	}

}