package main

type Request struct {
	Headers map[string]string
	Body string
	SSL bool

	Packet string	
	Built bool
}

func createRequest() *Request {
	return &Request{Headers: make(map[string]string)}
}

func (request *Request) Set(key, value string) {
	request.Headers["key"] = value
}

func (request *Request) SetHost(host string) {
	request.Headers["Host"] = host
}

func (request *Request) SetURI(URI string) {
	request.Headers["URI"] = URI
	request.SSL = getScheme(URI) == "https://"
}

func (request *Request) SetMethod(method string) {
	request.Headers["Method"] = method
}

func (request *Request) SetContentType(contentType string) {
	request.Headers["Content-Type"] = contentType
}

func (request *Request) SetCookie(cookie string) {
	request.Headers["Cookie"] = cookie
}

func (request *Request) SetBody(body string) {
	request.Body = body
}