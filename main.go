package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Start server on http://localhost:9999")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world!")
	})
	http.ListenAndServe(":9999", nil)
}
