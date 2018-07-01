package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello, World. Using official docker image")
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
