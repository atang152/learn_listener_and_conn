package netutil

import (
	"errors"
	"log"
	"net"
)

// Server defines the methods our server implementation must satisfy
type Server interface {
	Run() error
	Close() error
}

// TCPServer holds the basic structure of the TCP implementation
type TCPServer struct {
	addr   string
	server net.Listener
}

// NewServer creates a new Server using given protocol and addr
func NewServer(protocol, addr string) (Server, error) {

	if protocol == "tcp" {
		return &TCPServer{addr: addr}, nil
	}

	return nil, errors.New("invalid protocol provided")

}

// Run starts the TCP Server
func (t *TCPServer) Run() (err error) {

	// Listens to address and returns a listener interface which holds
	// the following method designs: Accept(), Close() and Addr()
	t.server, err = net.Listen("tcp", t.addr)
	if err != nil {
		log.Fatalln("failed to listent to address")
	}

	defer t.server.Close()
	log.Printf("server listening on %s\n", t.addr)

	for {
		// Wait for a connection and accepts it
		// conn implements reader and writer
		conn, err := t.server.Accept()

		if err != nil {
			log.Fatalln("could not accept connection", err)
			break
		}

		// Handle connection in a goroutine
		go t.handle(conn)

		defer conn.Close()
	}

	return
}

// Close shuts down the TCP Server
func (t *TCPServer) Close() (err error) {
	return t.server.Close()
}

// handle the connection in a new go routine
func (t *TCPServer) handle(conn net.Conn) {
	var buf = make([]byte, defaultBufSize)
	for {

		// Reads n bytes from client
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Fatalln("error reading from conn", err)
		}
		log.Printf("Read %d bytes from client, content is: %s\n", n, string(buf[:n]))

		// Write back n bytes to client
		n, err = conn.Write(buf[:n])
		if err != nil {
			log.Fatalln("error writing to conn", err)
		}

		log.Printf("Write %d bytes to client, content is: %s\n", n, string(buf[:n]))
	}
}
