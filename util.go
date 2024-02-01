package main

import (
	"strings"
)

func getScheme(URL string) string {
	var index int = strings.Index(URL, "://")
	if index >= 4 {
		var scheme string = URL[:index]
		if scheme == "http" || scheme == "https" {
			return scheme
		}
	}
	panic("Missing URL Scheme")
}

func getHost(URL, scheme string) string {
	var noSchemeURL string = URL[len(scheme + "://"):]
	return noSchemeURL[:strings.Index(noSchemeURL, "/")]
}

func isTLS(URL string) bool {
	return getScheme(URL) == "https"
}
