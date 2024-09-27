package main

import (
	"strings"
)

func getScheme(URI string) string {
	if URI[:8] == "https://" {
		return URI[:8]
	} else if URI[:7] == "http://" {
		return URI[:7]
	}
	panic("Missing URL scheme")
}

func getHost(URI string) string {
	var host string = URI[len(getScheme(URI)):]
	var index int = strings.Index(host, "/")
	if index != -1 {
		host = host[:index]	
	}
	return host
}

func getArguments(URI string) string {
	var index int = strings.Index(URI, "?")
	if index != -1 {
		return URI[index:]
	}
	return ""
}

func (request *Request) buldPacket() {
	if method, set := request.Headers["Method"]; set {
		request.Packet += method + " "
	} else {
		request.Packet += "GET " // We can assume they wanted to use a GET request
	}

	if URI, set := request.Headers["URI"]; set {
		request.Packet += URI + " HTTP/1.1\r\n"
	} else {
		panic("No URI supplied")
	}

	if host, set := request.Headers["Host"]; set {
		request.Packet += "Host: " + host + "\r\n"
	} else {
		request.Packet += "Host: " + getHost(request.Headers["URI"]) + "\r\n"
	}

	for key, value := range request.Headers {
		if key != "URI" && key != "Host" {
			request.Packet += key + ": " + value + "\r\n"
		}
	}
	request.Packet += "\r\n"

	if request.Body != "" {
		request.Packet += request.Body
	}
}