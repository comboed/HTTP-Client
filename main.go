package main

import "fmt"

func main() {
	var client *Client = createClient(1, 100, nil)
	var request *Request = createRequest()
	
	// Not SSL
	request.SetURI("http://example.org/")
	var body = client.Do(request)
	fmt.Println(string(body))

	// SSL
	request.SetURI("https://example.org/")
	body = client.Do(request)
	fmt.Println(string(body))
}