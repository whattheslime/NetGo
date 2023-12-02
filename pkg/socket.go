package netgo

import (
    "bufio"
    "crypto/tls"
	"crypto/x509"
    "encoding/base64"
    "fmt"
    "log"
    "net"
    "net/http"
    "net/url"
    "time"
)


// Socket Client
func (nObj NetObject) RunClient(
    cmd string, recv bool, send bool, proxy url.URL, isTls bool, cert string, 
    key string) {
    var conn net.Conn

    username := proxy.User.Username()
    password, _ := proxy.User.Password()

    if proxy.Scheme == "http" {
        // Step 1: Connect to the proxy server
        proxyConn, err := net.Dial(nObj.Type, proxy.Host)
        if err != nil {
            log.Fatalln("Error connecting to proxy server:", err)
            return
        }
        defer proxyConn.Close()

        // Step 2: Create a HTTP CONNECT request with handleing authentication
        request := fmt.Sprintf(
            "CONNECT %s HTTP/1.1\r\nHost: %s\r\n\r\n", 
            nObj.Service, nObj.Service)

        if username != "" && password != "" {
            credentials := base64.StdEncoding.EncodeToString(
                []byte(username+":"+password))
            request = fmt.Sprintf(
                "CONNECT %s HTTP/1.1\r\nHost: %s\r\n" + 
                "Proxy-Authorization: Basic %s\r\n\r\n", 
                nObj.Service, nObj.Service, credentials)
        }
        
        // Step 3: Send an HTTP CONNECT request to establish a tunnel with 
        // proxy authentication
        _, err = proxyConn.Write([]byte(request))
        if err != nil {
            log.Fatalln("Error sending CONNECT request to proxy:", err)
            return
        }

        // Step 4: Read the response from the proxy
        response, err := http.ReadResponse(bufio.NewReader(proxyConn), nil)
        if err != nil {
            log.Fatalln("Error reading CONNECT response from proxy:", err)
            return
        }

        if response.StatusCode != http.StatusOK {
            log.Fatalf(
                "Failed to establish tunnel. Proxy responded with " +
                "status code %d\n", response.StatusCode)
            return
        }

        log.Printf(
            "Connection established through proxy %s to %s\n", 
            proxy.Host, nObj.Service)

        conn = proxyConn
    // } else if proxy.Scheme == "socks5" || proxy.Scheme == "socks5h" {
    //     var auth = proxy.Auth
    //     auth.User = username
    //     auth.User = password
        
    //     conn, err := proxy.SOCKS5(nObj.Type, proxyauth)
    } else {
        // Try connection
        directConn, err := net.Dial(nObj.Type, nObj.Service)
        if err != nil {
            log.Fatalf("Connection refused")
        }
        
        conn = directConn
    }

    log.Println("Connected to", conn.RemoteAddr())

    // Add a TLS layer
    if isTls {
        tlsConn := nObj.TLSClient(conn, cert, key)
        handleConn(tlsConn, cmd, recv, send)
    } else {
        handleConn(conn, cmd, recv, send)
    }
}


// Socket Server
func (nObj NetObject) RunServer(
    cmd string, keepOpen bool, maxConns int, recv bool, send bool, 
    isTls bool, cert string, key string) {
    var listener net.Listener
    // Start listening
    if isTls {
        tlsListener := nObj.TLSServer(cert, key)
        listener = tlsListener
    } else {
        rawListener, err := net.Listen(nObj.Type, nObj.Service)
        if err != nil {
            log.Fatal(err)
        }
        listener = rawListener
    }
    
    log.Printf("Listening on %s://%s", nObj.Type, nObj.Service)

    for {
        // Wait for connection
        conn, err := listener.Accept()
        if err != nil {
            log.Fatalln("Connection failed:", err)
            continue
        }
        defer listener.Close()
        
        time.Sleep(1000)

        if len(clients) >= maxConns {
            conn.Close()
            log.Println(
                "New connection denied: connection limit reached: (" + 
                fmt.Sprint(maxConns) + ")")
            continue
        }

        log.Println("Connection from", conn.RemoteAddr())

        if isTls {
            // Checking connection state
            tlsConn, ok := conn.(*tls.Conn)
            if ok {
                log.Print("ok=true")
                state := tlsConn.ConnectionState()
                for _, v := range state.PeerCertificates {
                    log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
                }
            }
        }

        if keepOpen {
            go handleConn(conn, cmd, recv, send)
        } else {
            // Handle connection
            handleConn(conn, cmd, recv, send)
            // Close connection
            listener.Close()
            break
        }
    }
}
