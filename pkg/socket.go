package netgo

import (
	"log"
	"net"
)

// Socket Client
func (nObj NetObject) RunClient(cmd string) {
	// Try connection
	conn, err := net.Dial(nObj.Type, nObj.Service)
	if err != nil {
		log.Println("Connection refused")
		return
	}
	defer conn.Close()
	log.Println("Connected to", conn.RemoteAddr())

	// Handle connection
	handleConn(conn, cmd)
	log.Println("Broken pipe")
}

// Socket Server
func (nObj NetObject) RunServer(cmd string) {
	// Start listening
	listener, err := net.Listen(nObj.Type, nObj.Service)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on", nObj.Service, "...")
	defer listener.Close()

	// Wait for connection
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalln("Connection failed:", err)
	}
	defer conn.Close()
	log.Println("Connection receive from", conn.RemoteAddr())
	listener.Close()

	// Handle connection
	handleConn(conn, cmd)
	log.Println("Connection with", conn.RemoteAddr(), "closed")
}
