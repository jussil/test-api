package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, hostname)
}

func handlerDump(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, string(requestDump))

	for _, pair := range os.Environ() {
		fmt.Fprintln(w, pair)
	}
}

func main() {
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/dump", handlerDump)
	fmt.Println("Listening for requests..")
	http.ListenAndServe(":8080", nil)
}
