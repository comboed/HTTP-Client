package main

type Request struct {
	Packet map[string]string

}
func createRequest() *Request {
	return &Request {
		Packet: make(map[string]string),
	}
}

func (request *Request) Set(header, value string) {
	request.Packet[header] = value
}

func (request *Request) SetBody(body string) {
	request.Packet["Body"] = body
}

func (request *Request) DeleteHeader(header string) {
	delete(request.Packet, header)
}