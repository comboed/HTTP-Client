package main

import "fmt"

func main() {
	var headers map[string]string = make(map[string]string)
	headers["Method"] = "POST"
	headers["URL"] = "https://www.instagram.com/"
	fmt.Println(parseDefaultHeaders(headers))

	// URL := "http://www.xvideos.com/niggasex/lol/"
	// fmt.Println(getHost(URL, getScheme(URL)))
}