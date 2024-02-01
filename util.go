package main

type Request struct {
	Packet string
	Headers map[string]string
	SetHeaders []string

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

func (request *Request) Set(key, value string) {
	request.Headers[key] = value
}

func (request *Request) DeleteHeader(header string) {
	delete(request.Headers, header)
}

func (request *Request) IsHeaderSet(key string) {
	
}

func parseHeaders(request *Request) {
	if method := request.Headers["Method"]; method != "" {
		request.Packet += method
	} else {
		request.Headers["Method"] = "GET" // assume it was a GET request
	}
	if URL := request.Headers["URL"]; URL != "" {
		request.Packet += " " + URL + " HTTP/1.1\r\n"
	} else {
		panic("No URL supplied")
	}
}

