package main

import (
	"html"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.OpenFile("./log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "", log.LstdFlags)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Pong")
		logger.Printf("- %s - %s - %s - %d", html.EscapeString(r.Method), html.EscapeString(r.URL.Path), r.RemoteAddr, http.StatusOK)
	})
	http.HandleFunc("/bad-gateway", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, "Bad Gateway")
		logger.Printf("- %s - %s - %s - %d", html.EscapeString(r.Method), html.EscapeString(r.URL.Path), r.RemoteAddr, http.StatusBadGateway)
	})
	http.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMovedPermanently)
		io.WriteString(w, "Moved Permanently")
		logger.Printf("- %s - %s - %s - %d", html.EscapeString(r.Method), html.EscapeString(r.URL.Path), r.RemoteAddr, http.StatusMovedPermanently)
	})
	log.Fatalln(http.ListenAndServe(":9090", nil))
}
