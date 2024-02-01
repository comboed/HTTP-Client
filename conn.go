package main

import (
	"crypto/tls"
	"net"
)

func createConnection(address string) net.Conn {
	var conn, _ = net.Dial("tcp", address)
	return conn
}

func createTLSClient(conn net.Conn) tls.Conn {
	return *tls.Client(conn, &tls.Config{})
}

