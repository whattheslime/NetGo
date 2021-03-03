package conn

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"log"
)

// TLS Client
func (ntObj NetObject) RunTLSClient(cmd string) {
	// SSL Configuration
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert},
		InsecureSkipVerify: true}

	// Try connection
	conn, err := tls.Dial(ntObj.Type, ntObj.Service, &config)
	if err != nil {
		log.Fatalln("Connection failed:", err)
	}
	defer conn.Close()
	log.Println("Connected to", conn.RemoteAddr())

	// Checking connection state
	state := conn.ConnectionState()
	for _, v := range state.PeerCertificates {
		// fmt.Println(x509.MarshalPKIXPublicKey(v.PublicKey))
		log.Println("Issuer:", v.Issuer)
		log.Println("Subject:", v.Subject)
	}
	log.Println("Handshake complete: ", state.HandshakeComplete)
	log.Println("Protocol negotiation done: ",
		state.NegotiatedProtocolIsMutual)

	// Handle connection
	handleConn(conn, cmd)
	log.Println("Broken pipe")
}

// TLS Server
func (ntObj NetObject) RunTLSServer(cmd string) {
	// SSL Configuration
	cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader

	// Start listening
	listener, err := tls.Listen(ntObj.Type, ntObj.Service, &config)
	if err != nil {
		log.Fatal("Binding error:", err)
	}
	log.Println("Listening on", ntObj.Service, "...")

	// Wait for connection
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalln("Connection failed:", err)
	}
	defer conn.Close()
	log.Println("Connection receive from", conn.RemoteAddr())
	listener.Close()

	// Checking connection state
	tlsconn, ok := conn.(*tls.Conn)
	if ok {
		log.Print("ok=true")
		state := tlsconn.ConnectionState()
		for _, v := range state.PeerCertificates {
			log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
		}
	}

	// Handle connection
	handleConn(conn, cmd)
	log.Println("Connection with", conn.RemoteAddr(), "closed")
}
