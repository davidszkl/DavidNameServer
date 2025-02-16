package network

import (
	"fmt"
	"net"
)

func StartDns() {
	udpListener, err := net.ListenPacket("udp", "127.0.0.1:4269")
	if err != nil {
		fmt.Errorf("Error opening a socket.")
	}
	tcpListener, err := net.Listen("tcp", "127.0.0.1:4270")
	if err != nil {
		fmt.Errorf("Error opening a socket.")
	}
	defer udpListener.Close()
	defer tcpListener.Close()
	go listenToTcpConnection(tcpListener)
	listenToUdpConnection(udpListener)
}

func listenToTcpConnection(listener net.Listener) {
	for {
		tcpConn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("Error opening a TCP connection.")
		}

		go handleTcpRequest(tcpConn)
	}
}

func handleTcpRequest(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New client connected:", conn.RemoteAddr())

	// Read and write data to the client
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		}
		fmt.Printf("Received from %s: %s\n", conn.RemoteAddr(), string(buf[:n]))
	}
}

func listenToUdpConnection(listener net.PacketConn) {
	for {
		buffer := make([]byte, 4196)
		_, address, err := listener.ReadFrom(buffer)
		if err != nil {
			fmt.Errorf("Error reading UDP request.")
		}

		go handleUdpRequest(buffer, address)
	}
}

func handleUdpRequest(buffer []byte, address net.Addr) {
	fmt.Print(buffer, address.String())
}
