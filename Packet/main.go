package main

import (
	"fmt"
	"net"
	"time"
)

// -------------------------- 发包 --------------------------

func sendTCPPacket(target string, data string) bool {
	conn, err := net.DialTimeout("tcp", target, 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close() 

	_, err = conn.Write([]byte(data))
	if err != nil {
		return false
	}

	return true
}

// -------------------------- 抓包 --------------------------
func captureTCPPacket(listenAddr string) bool {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return false
	}
	defer listener.Close()

	fmt.Printf("开始监听TCP端口 %s\n", listenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleTCPConnection(conn)
	}

	return true
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		return
	}

	fmt.Printf("从 %s 抓取到数据: %s (长度: %d 字节)\n", clientAddr, string(buf[:n]), n)
}

func main() {
	
}
