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

func getHost(URL string) string {
	var noSchemeURL string = URL[len(getScheme(URL) + "://"):]
	var slashIndex, queryIndex int = strings.Index(noSchemeURL, "/"), strings.Index(noSchemeURL, "?")
	if slashIndex != -1 {
		return noSchemeURL[:slashIndex]
	} else if queryIndex != -1 {
		return noSchemeURL[:queryIndex]
	}
	return noSchemeURL
}

func isTLS(URL string) bool {
	return getScheme(URL) == "https"
}
