package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type Jar struct {
	jar map[string][]*http.Cookie
}

func (p *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	fmt.Printf("Cookie set: %s\n", cookies)
	p.jar[u.Host] = cookies
}

func (p *Jar) Cookies(u *url.URL) []*http.Cookie {
	fmt.Printf("Cookie being returned is : %s\n", p.jar[u.Host])
	return p.jar[u.Host]
}
