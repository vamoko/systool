package main

import (
	"fmt"
	"net/url"
)

func main() {
	urls := []string{
		"https",
		"https://",
		"",
		"www.bing.com",
		"www.163.com/",
		"http://www",
		"http://www.dumpsters.com",
		"http://www.baidu.com/",
		"http://www.baidu.com/123",
		"http://www.baidu.com/123/",
		"https://www.dumpsters.com:443",
		"/testing-path",
		"testing-path",
		"alskjff#?asf//dfas",
	}
	for _, u := range urls {
		val, err := url.Parse(u)
		scheme := val.Scheme
		host := val.Host
		hostname := val.Hostname()
		path := val.Path

		fmt.Println("===[ " + u + " ]===")
		fmt.Println("val        :", val)
		fmt.Println("error      :", err)
		fmt.Println("scheme     :", scheme)
		fmt.Println("host       :", host)
		fmt.Println("hostname   :", hostname)
		fmt.Println("path       :", path)
		fmt.Println()

		uu, err := url.ParseRequestURI(u)
		fmt.Println("uri        :", uu)
		fmt.Println("error      :", err)
		fmt.Println()
	}
}
