package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/prashantv/go_profiling_talk/handlers"
)

const hostPort = ":9090"

func main() {
	flag.Parse()

	// Register our two handlers, and an index page that links to these handlers.
	// Register 2 versions of the Hello handler:
	// 1 with the stats profiling.
	// 1 without the stats.
	http.HandleFunc("/hello", handlers.WithStats(handlers.Hello))
	http.HandleFunc("/simple", handlers.Hello)
	http.HandleFunc("/", index)

	fmt.Println("Starting server on", hostPort)
	if err := http.ListenAndServe(hostPort, nil); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}

func index(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "text/html")
	io.WriteString(w, "<h2>Links</h2>\n<ul>")
	for _, link := range []string{"/hello", "/simple"} {
		fmt.Fprintf(w, `<li><a href="%v">%v</a>`, link, link)
	}
	io.WriteString(w, "</ul>")
}
