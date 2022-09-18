package main

import (
	"fmt"
	"net/http"
	"flag"
)
var port = flag.Int("p", 8080, "server port")

func main() {
	flag.Parse()
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/mo", HelloServer1)
	http.HandleFunc("/sherr", HelloServer2)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func HelloServer1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Mo %s!", r.URL.Path[1:])
}

func HelloServer2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Sherr %s!", r.URL.Path[1:])
}
