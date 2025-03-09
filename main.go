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
)

func getServerAddr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func handler(addr string, dir string) http.Handler {
	fmt.Printf("Listening on http://%s\n", addr)
	return http.FileServer(http.Dir(dir))
}

func main() {
	flag.Parse()

	addr := getServerAddr(*host, *port)
	if err := http.ListenAndServe(addr, handler(addr, *dir)); err != nil {
		panic(err)
	}
}
