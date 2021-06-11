package ltcp

import (
	"fmt"
	"net"
)

func TcpServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err", err)
			return
		}
		go TcpConnHandler(conn)
	}

}

func TcpConnHandler(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr()
	fmt.Println("addr:", addr)
	buf := make([]byte, 4096)
	for {
		read, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err", err)
			return
		}
		fmt.Println("server read ", read)
	}
}
