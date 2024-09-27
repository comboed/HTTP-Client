package main

import (
	"github.com/valyala/fasthttp"
	"crypto/tls"
	"bufio"
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
		Pool: make(chan *Connection, maxConnections + 100),
	}
}

func (client *Client) GetConnection(request *Request, reDial bool) *Connection {	
	if len(client.Pool) < client.MaxConnections {
		var connection *Connection = createConnection(client.TLSConfig)
		connection.DialHost(request)
		return connection
	}
	return <- client.Pool
}

func (client *Client) Do(request *Request) []byte {
	var connection *Connection = client.GetConnection(request, false)
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	for {
		if _, err := connection.Conn.Write([]byte(request.buldPacket())); err != nil {
			connection.Conn.Close()
			connection = client.GetConnection(request, true)
			continue
		}
		if err := response.Read(bufio.NewReader(connection.Conn)); err != nil {
			continue
		}
		client.Pool <- connection
		return response.Body()
	}
}
