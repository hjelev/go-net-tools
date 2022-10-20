# go net tools


# check_ssl 

check the issuer and the expiration date of a domain certificate

## Example:

```
$ ./check_ssl masoko.net
Issuer: CN=R3,O=Let's Encrypt,C=US
Expiry: Monday, 02-Jan-23 19:48:21 UTC
```

# check_redirect

check destination url and status code of url

## Example:

```
$ ./ceck_redirect google.com 
301 https://www.google.com/
```