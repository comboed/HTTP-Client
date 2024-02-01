package main

import (
	"strings"
)

func getScheme(URL string) string {
	var index int = strings.Index(URL, "://")
	if index == 4 || index == 5 {
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

func parseDefaultHeaders(headers map[string]string) string {
	var packet string
	if method := headers["Method"]; method != "" {
		packet = method
	} else {
		packet = "GET" // assume it was a GET request
	}
	if URL := headers["URL"]; URL != "" {
		packet += " " + URL + " HTTP/1.1\r\n"
	} else {
		panic("No URL Suppllied")
	}
	return packet
}

func parseExtraHeaders() {

}

// func getMethod(headers map [string]string) string {
	
// }