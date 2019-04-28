package netutil

import (
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Initialize a new TCP server
func init() {
	// Start the new server
	svr, err := NewServer("tcp", ":1234")
	if err != nil {
		log.Fatalln("error starting TCP server", err)
	}

	// Run the servers in goroutines to stop blocking
	go func() {
		err = svr.Run()
		if err != nil {
			log.Fatalln("error running TCP server", err)
		}
	}()

	time.Sleep(time.Second * 1)
}

// TestNewServer simply checks whether a server is up and accepting connection
func TestNewServer(t *testing.T) {

	conn, err := net.Dial("tcp", ":1234")

	assert.Nil(t, err)

	defer conn.Close()
}
