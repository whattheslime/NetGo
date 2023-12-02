package netgo

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"log"
	"time"
	"net"
)


func genCert() ([]byte, []byte){
	// Generate a private key
    key, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        log.Fatal("Private key cannot be created.", err.Error())
    }
    // Generate a pem block with the private key
	keyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

    // Dump private key to bytes
    tml := x509.Certificate{
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(5, 0, 0),
		SerialNumber: big.NewInt(123123),
		Subject: pkix.Name{
			CommonName:   "New Name",
			Organization: []string{"New Org."},
		},
		BasicConstraintsValid: true,
	}
	cert, err := x509.CreateCertificate(
		rand.Reader, &tml, &tml, &key.PublicKey, key)
	if err != nil {
		log.Fatalf("Certificate cannot be created: %s", err)
	}
	
	// Generate a pem block with the certificate
	certPem := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})
	return certPem, keyPem
}

func loadTLSConfig(cert string, key string) (tls.Certificate, error) {
	var tlsCert tls.Certificate
	var err error

	if cert == "" && key == "" {
		// Genrerating Certificate
		cert, key := genCert()
		tlsCert, err = tls.X509KeyPair(cert, key)
	} else {
		// Load certificate and key from files
		tlsCert, err = tls.LoadX509KeyPair(cert, key)
	}
	return tlsCert, err
}

// TLS Client
func (nObj NetObject) TLSClient(conn net.Conn, cert string, key string) (net.Conn) {
	// TLS Configuration
	tlsCert, err := loadTLSConfig(cert, key)
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{tlsCert},
		InsecureSkipVerify: true}

	// Try connection
	tlsConn := tls.Client(conn, &config)

	// Checking connection state
	state := tlsConn.ConnectionState()
	for _, v := range state.PeerCertificates {
		// fmt.Println(x509.MarshalPKIXPublicKey(v.PublicKey))
		log.Println("Issuer:", v.Issuer)
		log.Println("Subject:", v.Subject)
	}
	
	log.Println("Handshake complete: ", state.HandshakeComplete)
	log.Println("Protocol negotiation done: ",
		state.NegotiatedProtocolIsMutual)

	return tlsConn
}


func (nObj NetObject) TLSServer(cert string, key string) net.Listener {
	// TLS Configuration
	tlsCert, err := loadTLSConfig(cert, key)
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{tlsCert}}
	config.Rand = rand.Reader

	// Start listening
	listener, err := tls.Listen(nObj.Type, nObj.Service, &config)
	if err != nil {
		log.Fatal("Binding error:", err)
	}

	return listener
}
