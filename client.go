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
		Pool: make(chan *Connection, maxConnections),
	}
}

func (client *Client) GetConnection(request *Request, reDial bool) *Connection {
	if len(client.Pool) < client.MaxConnections || reDial {
		var connection *Connection = createConnection(client.TLSConfig)
		connection.DialHost(request)
		return connection
	} else {
		return <- client.Pool
	}
}

func (client *Client) Do(request *Request) []byte {
	var connection *Connection = client.GetConnection(request, false)
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	request.buldPacket() // temp for now, figure out a new way to not constantly build packets on Do
	for {
		if _, err := connection.Conn.Write([]byte(request.Packet)); err != nil {
			connection = client.GetConnection(request, true)
			continue
		}
		client.Pool <- connection
		response.Read(bufio.NewReader(connection.Conn))
		return response.Body()
	}
}