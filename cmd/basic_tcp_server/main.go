package main

import (
	"log"
	"os"
	"time"

	"github.com/atang152/learn_listener_and_conn/pkg/netutil"
)

func main() {

	// Check input argument
	if len(os.Args) <= 1 {
		log.Println("usage: go run main.go YOUR_STRING")
	}

	// Start the new server.
	srv, err := netutil.NewServer("tcp", ":1234")
	if err != nil {
		log.Fatalln("error starting TCP server", err)
	}

	// Run the server in Goroutine
	go func() {
		err = srv.Run()
		if err != nil {
			log.Fatalln("error running TCP server", err)
		}
	}()

	// Use time sleep to block before running client
	time.Sleep(time.Second * 1)

	msg := os.Args[1]

	// Initiate a new client
	client := netutil.TCPClient{Addr: ":1234", Msg: msg}

	err = client.Run()
	if err != nil {
		log.Fatalln("error encountered while running client", err)
	}

}
