package main

import (
	"canvas-server/di"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.DefaultServeMux

	di.ResolveSubscriber()(mux)
	di.ResolveAPI()(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("running server on port: %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("failed running server, err=%+v", err)
	}
}
