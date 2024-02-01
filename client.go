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
		Connections: []*tls.Conn{},
		ConnectionPool: make(chan []*tls.Conn),
	}
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

func fillConnectionPool(connections *[]*tls.Conn) {
	//appendConnection(connections, "")
}
