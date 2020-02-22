package gui

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

func GuiServer() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":8080", nil)
}
