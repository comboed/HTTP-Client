package main

import (
	"crypto/tls"
	_"net"
)

type Client struct {
	ReadBufferSize int
	MaxConnections int
	Connections []*tls.Conn
	ConnectionPool chan []*tls.Conn
}

func createClient(readBufferSize, MaxConnections int) *Client {
	return &Client {
		ReadBufferSize: readBufferSize,
		MaxConnections: MaxConnections,
		Connections: createConnectionSlice(MaxConnections),
		ConnectionPool: make(chan []*tls.Conn),
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

func fillChannelPool(connections *[]*tls.Conn, connectionPool chan *tls.Conn) {
	for {
		for _, connection := range *connections {
			connectionPool <- connection
		}
	}
}
