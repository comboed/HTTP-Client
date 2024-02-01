package main

import (
	"crypto/tls"
	_"net"
)

type Client struct {
	ReadBufferSize int
	MaxConnections int
	Connections []*tls.Conn
	connectionIndex chan int
}

func createClient(readBufferSize, MaxConnections int) *Client {
	return &Client {
		ReadBufferSize: readBufferSize,
		MaxConnections: MaxConnections,
		Connections: createConnectionSlice(MaxConnections),
		connectionIndex: make(chan int),
	}
}

func createConnectionSlice(maxConnections int) (connections []*tls.Conn) {
	for i := 0; i < maxConnections; i++ {
		connections = append(connections, nil)
	}
	return connections
}

func createConnection(address string) *tls.Conn {
	var conn, _ = tls.Dial("tcp", address, &tls.Config{})
	return conn
}

func appendConnection(connections *[]*tls.Conn, URL string) {
	if isTLS(URL) {
		*connections = append(*connections, createConnection(getHost(URL) + ":443"))
	} else {
		*connections = append(*connections, createConnection(getHost(URL) + ":80"))
	}
}

func fillChannelPool(connections *[]*tls.Conn, connectionIndex chan int) {
	for {
		for i := range *connections {
			connectionIndex <- i
		}
	}
}


func (client *Client) Do(request *Request) {
	var conn *tls.
}
