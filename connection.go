package main

import (
	"crypto/tls"
	"net"
)

type Connection struct {   
	Conn net.Conn  
	TLSConfig *tls.Config
}

func createConnection(tlsConfig *tls.Config) *Connection {
	return &Connection{TLSConfig: tlsConfig}
}

func (connection *Connection) DialHost(request *Request) {
	var host string = getHost(request.Headers["URI"])
	if getScheme(request.Headers["URI"]) == "http://" {
		connection.Conn, _ = net.Dial("tcp", host + ":80")
	} else {
		connection.Conn, _ = tls.Dial("tcp", host + ":443", connection.TLSConfig)
	}
}