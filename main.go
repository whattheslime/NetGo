package main

import (
    "fmt"
    "os"
    "strconv"
    "net/url"
    
    "github.com/WhatTheSlime/NetGo/pkg"
    "github.com/spf13/pflag"
)


var (
    Command string
    Hostname string
    KeepOpen bool
    Listen bool
    MaxConns int
    Port int
    Proxy string
    ProxyUrl url.URL
    RecvOnly bool
    SendOnly bool
    TLS bool
    TLSCert string
    TLSKey string
    Verbose bool
    Version bool
    UDP bool
)


const DefaulPort = 31337

// Check port validity.
func checkPort(port string) (int, error) {
    intPort, err := strconv.Atoi(port)
    if err != nil || intPort < 1 || intPort > 65536 {
        return intPort, fmt.Errorf("Invalid port number \"%d\".", Port)
    }
    return intPort, nil
}


// Check arguments validity.
func checkArgs(args []string) error {
    // Check for displaying version only
    if Version {
        return fmt.Errorf("NetGo Version %s (%s)", netgo.VERSION, netgo.GITHUB)
    }

    // Check host and ports arguments
    switch {
        case len(args) == 2:
            Hostname = args[0]
            intPort, err := checkPort(args[1]);
            if err != nil {return err} else {Port = intPort}
        case len(args) == 1 && Listen:
            intPort, err := checkPort(args[0]); 
            if err != nil {
                Hostname = args[0]
                Port = DefaulPort
            } else {
                Hostname = "0.0.0.0"
                Port = intPort
            }
        case len(args) == 1:
            Hostname = args[0]
            Port = DefaulPort
        case len(args) == 0 && Listen:
            Hostname = "0.0.0.0"
            Port = DefaulPort
        case len(args) == 0:
            return fmt.Errorf("You must specify a host to connect to.")
        default:
            return fmt.Errorf("Too many arguments.")
    }

    // Check non listening options
    if ! Listen {
        if KeepOpen {
            return fmt.Errorf("--listen option is mandatory for --keep-open.")
        }
        if Proxy != "" {
            url, err := url.Parse(Proxy); if err != nil {
                return fmt.Errorf("Proxy bad format: %s\n", err)
            }
            ProxyUrl = *url
        }

    } else {
        if Proxy != "" {
            return fmt.Errorf("--listen cannot be used with --proxy.")
        }
    }

    // Check TLS args
    if (len(TLSCert) > 0) != (len(TLSKey) > 0) {
        return fmt.Errorf(
            "The --tls-key and --tls-cert options must be used together.")
    }
    return nil
}


// Add persistent flags
func initFlags() {
    pflag.StringVarP(
        &Command, "exec", "e", "", "Executes the given command")
    pflag.BoolVarP(
        &KeepOpen, "keep-open", "k", false, 
        "Accept multiple connections in listen mode")
    pflag.BoolVarP(
        &Listen, "listen", "l", false, 
        "Bind and listen for incoming connections")
    pflag.IntVarP(
        &MaxConns, "max-conns", "m", 50, 
        "Maximum simultaneous connections (default: 50)")
    pflag.StringVarP(
        &Proxy, "proxy", "x", "", 
        "Specify url of host proxy through " + 
        "http://<username>:<password>@<host>:<port>")
    pflag.BoolVarP(
        &RecvOnly, "recv", "", false, 
        "Only receive data, never send anything")
    pflag.BoolVarP(
        &SendOnly, "send", "", false, 
        "Only send data, ignoring received; quit on EOF")
    pflag.BoolVar(
        &TLS, "ssl", false, "Connect or listen with SSL/TLS")
    pflag.StringVar(
        &TLSCert, "ssl-cert", "", "Load SSL Certificate")
    pflag.StringVar(
        &TLSKey, "ssl-key", "", "Load SSL key")
    pflag.BoolVarP(
        &Verbose, "verbose", "v", false, "Set verbose output")
    pflag.BoolVar(
        &Version, "version", false, "Display version information and exit")
    // pflag.BoolVarP(
    //     &UDP, "udp", "u", false, "Use UDP instead of default TCP")
}


// Program entry point.
func main() {
    initFlags()
    
    pflag.ErrHelp = nil
    pflag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: netgo [options] [hostname] [port]\n\n")
		pflag.PrintDefaults()
        fmt.Fprintf(os.Stderr, "\n")
        os.Exit(1)
    }

    pflag.Parse()
    
    err := checkArgs(pflag.Args()); 
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    } else {
        protocol := "tcp"
        if UDP {
            protocol = "udp"
        }

        nObj := netgo.NetObject{
            Type:    protocol,
            Service: fmt.Sprintf("%s:%d", Hostname, Port),
        }

        if Listen {
            nObj.RunServer(
                Command, KeepOpen, MaxConns, RecvOnly, SendOnly, 
                TLS, TLSCert, TLSKey)
        } else {
            nObj.RunClient(
                Command, RecvOnly, SendOnly, ProxyUrl, 
                TLS, TLSCert, TLSKey)
        }
    }
}
