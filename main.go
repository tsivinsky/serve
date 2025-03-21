package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	dir  = flag.String("S", ".", "directory to serve")
	host = flag.String("H", "localhost", "server host")
	port = flag.Int("p", 5000, "server port")

	certFile    = flag.String("cert", "", "ssl certificate file")
	certKeyFile = flag.String("key", "", "ssl certificate key file")
)

func getServerAddr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func handler(addr string, dir string, https bool) http.Handler {
	protocol := "http:"
	if https {
		protocol = "https:"
	}

	fmt.Printf("Listening on %s//%s\n", protocol, addr)
	return http.FileServer(http.Dir(dir))
}

func main() {
	flag.Parse()

	var err error
	addr := getServerAddr(*host, *port)
	if *certFile != "" && *certKeyFile != "" {
		err = http.ListenAndServeTLS(addr, *certFile, *certKeyFile, handler(addr, *dir, true))
	} else {
		err = http.ListenAndServe(addr, handler(addr, *dir, false))
	}
	if err != nil {
		panic(err)
	}
}
