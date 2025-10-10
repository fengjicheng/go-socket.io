package main

import (
	"log"
	"net/http"

	"go-socket.io/engineio"
	"go-socket.io/logger"
)

func main() {
	s := engineio.NewServer(nil)
	defer s.Shutdown()
	http.Handle("/socket.io/", s)
	http.Handle("/", http.FileServer(http.Dir("./asset")))

	logger.Info("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
