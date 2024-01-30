package main

import (
	_ "fmt"
	"strings"
)

func getScheme(URL string) string {
	var index int = strings.Index(URL, "://")
	if index != 4 && index != 5 {
		panic("Missing URL Scheme")
	}
	return URL[:index]
}

func isTLS(URL string) bool {
	return getScheme(URL) == "https"
}

func getHost(URL string) {
	
}