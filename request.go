package main

type Request struct {
	Packet string
	Headers map[string]string
	Body string
}

func createRequest() *Request {
	return &Request {
		Headers: make(map[string]string),
	}
}

func (request *Request) SetMethod(method string) {
	request.Headers["Method"] = method
}

func (request *Request) SetURL(URL string) {
	request.Headers["URL"] = URL
}

func (request *Request) SetHost(host string) {
	request.Headers["Host"] = host
}

func (request *Request) SetBody(body string) {
	request.Headers["Body"] = body
}

func (request *Request) SetHeader(key, value string) {
	request.Headers[key] = value
}

func (request *Request) DeleteHeader(header string) {
	delete(request.Headers, header)
}

func buildPacket(request *Request) {
	if method, exists := request.Headers["Method"]; exists {
		request.Packet = method
	} else {
		request.Packet = "GET" // assume it was a GET request
	}
	if URL, exists := request.Headers["URL"]; exists {
		request.Packet += " " + URL + " HTTP/1.1\r\n"
	} else {
		panic("No URL supplied")
	}
	if host, exists := request.Headers["Host"]; exists {
		request.Packet += "Host: " + host + "\r\n"
	} else {
		request.Packet += "Host: " + getHost(request.Headers["URL"]) + "\r\n"
	}
	for key, value := range request.Headers {
		if key != "Method" && key != "URL" && key != "Host" {
			request.Packet += key + ": " + value + "\r\n"
		}
	}
	request.Packet += "\r\n"
}
