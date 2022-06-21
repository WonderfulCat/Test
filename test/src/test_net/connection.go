package test_net

import "net"

type ConnTest struct {
	Conn net.Conn
	Name string
}

func NewConnTest(conn net.Conn, name string) *ConnTest {
	return &ConnTest{Conn: conn, Name: name}
}
