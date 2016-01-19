package handlers

import (
	"fmt"
	"net/http"
)

// Hello is the "hello world" endpoint.
func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
