// Package main is the application entry point. It supports running as an HTTP server
// or as an interactive CLI based on flags.
package main

import (
	"GoApi/cmd/cli"
	"GoApi/cmd/server"
	"flag"
)

// main parses command-line flags and starts either the HTTP server or the CLI menu.
// Use -server to start the server, -cli to start the CLI; if no flag is set, the server runs by default.
func main() {
	runAsServer := flag.Bool("server", false, "Start HTTP Server")
	runAsCli := flag.Bool("cli", false, "Start CLI Menu")
	flag.Parse()

	if *runAsServer {
		server.RunServer()
		return
	}
	if *runAsCli {
		cli.UserCliMenu()
		return
	}
	server.RunServer()
}
