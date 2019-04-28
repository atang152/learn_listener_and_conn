package netutil

import (
	"log"
	"net"
)

// Client defines the methods our client implementation must satisfy
type Client interface {
	Run() error
	Close() error
}

// TCPClient holds an address and a net.Conn
type TCPClient struct {
	Addr string
	Msg  string
}

// Run dials to a connection on a given a given address
func (c *TCPClient) Run() (err error) {

	// Dial connects to the address on the named network and returns a connection
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		log.Fatalln("dial error", err)
	}

	defer conn.Close()
	log.Printf("client successfully dialed to %s\n", c.Addr)

	clientHandler(conn, c.Msg)

	return
}

// ClientHandler writes bytes to a server once and reads back the server's response
func clientHandler(conn net.Conn, msg string) {

	// Create a buffer with 4096 bytes
	var buf = make([]byte, defaultBufSize)

	// Writes to server
	n, err := conn.Write([]byte(msg))
	if err != nil {
		log.Fatalln("Error writing from conn...", err)
	}

	log.Printf("Write %d bytes to server, content is: %s\n", n, string(msg))

	// Read from server
	n, err = conn.Read(buf[:])
	if err != nil {
		log.Fatalln("Error reading from conn...", err)
	}

	log.Printf("Read %d bytes from server, content is: %s\n", n, string(buf[:n]))

	// Wait forever
	select {}
}
