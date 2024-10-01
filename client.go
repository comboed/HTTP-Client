package main

import (
	"github.com/valyala/fasthttp"
	"crypto/tls"
	"bufio"
	"net"
)

type Client struct {
	MaxConnections int
	ReadBufferSize int
	TLSConfig *tls.Config
	Pool chan *Connection
}

func createClient(maxConnections, readBufferSize int, tlsConfig *tls.Config) *Client {
	return &Client{
		MaxConnections: maxConnections,
		ReadBufferSize: readBufferSize,
		TLSConfig: tlsConfig,
		Pool: make(chan *Connection, maxConnections),
	}
}

func (client *Client) GetConnection(request *Request, reDial bool) *Connection {	
	if len(client.Pool) < client.MaxConnections || reDial {
		var connection *Connection = createConnection(client.TLSConfig)
		connection.DialHost(request)
		return connection
	}
	return <- client.Pool
}

func (client *Client) ReadResponse(connection net.Conn) ([]byte, error) {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	var err error
	defer fasthttp.ReleaseResponse(response)
	
	if client.ReadBufferSize > 0 {
		err = response.ReadLimitBody(bufio.NewReader(connection), client.ReadBufferSize)
	} else {
		err = response.Read(bufio.NewReader(connection))
	}
	return response.Body(), err
}

func (client *Client) Do(request *Request) []byte {
	var connection *Connection = client.GetConnection(request, false)
	for i := 0; i < 10; i++ {
		if _, err := connection.Conn.Write([]byte(request.buldPacket())); err != nil {
			connection.Conn.Close()
			connection = client.GetConnection(request, true)
			continue
		}

		var body, err = client.ReadResponse(connection.Conn)
		if err != nil {
			continue
		}
		client.Pool <- connection
		return body
	}
	panic("Failed to send request after 10 attempts")
}
