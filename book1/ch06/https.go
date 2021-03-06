package main

import (
	"fmt"
	"log"
	"net/http"
)

const SERVER_PORT = 8080
const SERVER_DOMAIN = "localhost"
const RESPONSE_TEMPLATE = "hello https"

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}

func main() {
	http.HandleFunc(fmt.Sprintf("%s%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "rui.crt", "rui.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
