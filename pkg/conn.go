package netgo

import (
    "fmt"
    "io"
    "net"
    "os"
    "log"
    "os/exec"
    "bufio"
    "sync"
    "strings"
)

type Client struct {
    Conn    net.Conn
    Writer  bufio.Writer
}

// Network object
type NetObject struct {
    Type    string
    Service string
}

var (
    clients         []Client
    clientsMutex    sync.Mutex
)

type customOutput struct{}

func (c customOutput) Write(p []byte) (int, error) {
    fmt.Println("received output: ", string(p))
    return len(p), nil
}

// Manage connection for different behavior
func handleConn(conn net.Conn, binPath string, recvOnly bool, sendOnly bool) {
    defer func() {
        conn.Close()
        clientsMutex.Lock()
        for i, client := range clients {
            if client.Conn == conn {
                clients = append(clients[:i], clients[i+1:]...)
            }
        }
        log.Printf(
            "Connection with %s closed (Connected clients: %d)", 
            conn.RemoteAddr(), len(clients))
        clientsMutex.Unlock()
    }()

    connReader := bufio.NewReader(conn)
    connWriter := bufio.NewWriter(conn)
    inReader := bufio.NewReader(os.Stdin)
    outWriter := bufio.NewWriter(os.Stdout)
    
    // Add client to buffer
    clientsMutex.Lock()
    clients = append(clients, Client{conn, *connWriter})
    log.Println("Connected sockets:", len(clients))  
    clientsMutex.Unlock()
    
    if sendOnly {
        // Send only mode
        SendData(inReader)
    } else if recvOnly {
        // Receive only mode
        RecvData(connReader, outWriter)
    } else if binPath != "" {
        // Execute command and send Standard I/O net.Conn
        args := strings.Fields(binPath)

        cmd := exec.Command(args[0], args[1:]...)

        // Start the command
        cmd.Stdin = conn
        cmd.Stdout = conn
        cmd.Stderr = outWriter

        err := cmd.Run(); if err != nil {
            log.Println("Error executing command:", err)
            return
        }
    } else {
        // Copy data from stdin to all remote clients
        go SendData(inReader)
    
        // Copy data from remote client to stdout
        RecvData(connReader, outWriter)
    }
    return
}


func SendData(inReader *bufio.Reader) {
    for {
        toSendData, errorData := inReader.ReadString('\n')
        clientsMutex.Lock()
        if errorData == io.EOF {
            break
        }
        for _, client := range clients {
            fmt.Fprint(client.Conn, toSendData)
        }
        clientsMutex.Unlock()
    }
}

func RecvData(connReader *bufio.Reader, outWriter *bufio.Writer) {
    for {
        receiveData, errorData := connReader.ReadString('\n')
        if errorData == io.EOF {
            break
        } else {
            fmt.Fprint(outWriter, receiveData)
            outWriter.Flush()
        }
    }
}

