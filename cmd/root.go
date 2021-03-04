package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/WhatTheSlime/netgo/pkg"
)


var (
	Command string
	Hostname string
	TLS bool
	TLSCert string
	TLSKey string
	UDP bool
	Listen bool
	Port string
	Verbose bool
	Version bool

	rootCmd = &cobra.Command{
		Use:   fmt.Sprintf("%s [hostname] [port]", os.Args[0]),
		Short: "A basic implementation of ncat in go language",
		Args:  checkArgs,
		Run: func(cmd *cobra.Command, args []string) {
			protocol := "tcp"
			if UDP {
				protocol = "udp"
			}

			nObj := netgo.NetObject{
				Type:    protocol,
				Service: fmt.Sprintf("%s:%s", Hostname, Port),
			}
			switch {
			case TLS && Listen:
				nObj.RunTLSServer(Command, TLSCert, TLSKey)
			case TLS:
				nObj.RunTLSClient(Command, TLSCert, TLSKey)
			case Listen:
				nObj.RunServer(Command)
			default:
				nObj.RunClient(Command)
			}
		},
	}
)

func checkArgs(cmd *cobra.Command, args []string) error {
	// Check for displaying version only
	if Version {
		log.Printf("Version %s ( %s )\n", netgo.VERSION, netgo.GITHUB)
		os.Exit(0)
	}
	switch {
	case len(args) == 2:
		Hostname = args[0]
		Port = args[1]
	case len(args) == 1 && Listen:
		Hostname = "0.0.0.0"
		Port = args[0]
	case len(args) == 0:
		return fmt.Errorf("Missing arguments")
	default:
		return fmt.Errorf("Too many arguments")
	}
	port, err := strconv.Atoi(Port)
	if err != nil || port < 1 || port > 65536 {
		return fmt.Errorf("Invalid port number %s", Port)
	}

	// Check TLS args
	if (len(TLSCert) > 0) != (len(TLSKey) > 0) {
		return fmt.Errorf(
			"The --tls-key and --tls-cert options must be used together")
	}
	return nil
}

func init() {
	// Add persistent flags
	rootCmd.PersistentFlags().StringVarP(
		&Command, "exec", "e", "", "Executes the given command")
	rootCmd.PersistentFlags().BoolVarP(
		&Verbose, "verbose", "v", false, "Set verbose output")
	rootCmd.PersistentFlags().BoolVarP(
		&Listen, "listen", "l", false, 
		"Bind and listen for incoming connections")
	// rootCmd.PersistentFlags().BoolVarP(
	//	&UDP, "udp", "u", false, "Use UDP instead of default TCP")
	rootCmd.PersistentFlags().BoolVar(
		&TLS, "tls", false, "Connect or listen with TLS")
	rootCmd.PersistentFlags().StringVar(
		&TLSCert, "tls-cert", "", "Load TLS Certificate")
	rootCmd.PersistentFlags().StringVar(
		&TLSKey, "tls-key", "", "Load TLS key")
	rootCmd.PersistentFlags().BoolVar(
		&Version, "version", false, "Display version information and exit")
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
