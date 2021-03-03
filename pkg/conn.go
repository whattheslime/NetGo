package conn

import (
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"os/exec"
)

// Network object
type NetObject struct {
	Type    string
	Service string
}

func connCloseHandler(conn net.Conn) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		conn.Close()
		os.Exit(0)
	}()
}

func handleConn(conn net.Conn, binPath string) {
	if binPath != "" {
		// Execute command and send Standard I/O net.Conn
		cmd := exec.Command(binPath)
		cmd.Stdin = conn
		cmd.Stdout = conn
		cmd.Stderr = conn
		cmd.Run()
	} else {
		// Copy Standard I/O in a net.Conn
		go io.Copy(os.Stderr, conn)
		go io.Copy(os.Stdout, conn)
		io.Copy(conn, os.Stdin)

	}
}
