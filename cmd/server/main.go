// Created by Petr Lozhkin

package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	panic(http.ListenAndServe(":8080", nil))
}
