package main

import (
	"flag"

	"github.com/freeloginname/home_work_basic/hw13_http/client"
	"github.com/freeloginname/home_work_basic/hw13_http/server"
)

func main() {
	mode := flag.String("mode", "server", "server or client mode")
	ip := flag.String("ip", "127.0.0.1", "IP address")
	port := flag.String("port", "8080", "Port number")
	url := flag.String("url", "http://127.0.0.1:8080", "server url")
	method := flag.String("method", "GET", "method")
	path := flag.String("path", "get_user", "path")
	flag.Parse()
	switch *mode {
	case "client":
		client.Client(url, method, path)
	case "server":
		server.Server(ip, port)
	default:
		panic("unknown mode")
	}
}
