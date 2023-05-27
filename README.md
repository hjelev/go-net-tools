# go net tools

# check_ssl 

Check the issuer and the expiration date of a domain certificate.

## Example:

```
$ ./check_ssl masoko.net
Issuer: CN=R3,O=Let's Encrypt,C=US
Expiry: Monday, 02-Jan-23 19:48:21 UTC
```

# check_redirect

Usage: check_redirect [URL] 

Display status code returned from url and follow redirects.

Automatically appends https:// to url if not present.\n

## Example:

```
$ ./check_redirect http://amazon.com                                                                       127 ↵ masoko@ubuntu-pc
301 http://amazon.com
301 https://amazon.com/
200 https://www.amazon.com/
```