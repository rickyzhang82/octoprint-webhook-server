package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rickyzhang82/octoprint-webhook-server/handler"
)

const (
	defaultPort    = "12000"
	portEnvVarName = "WEBHOOK_SEREVER_PORT"
)

func main() {
	http.HandleFunc("/done", handler.Done)
	portEnvVar := os.Getenv(portEnvVarName)
	port := defaultPort
	if 0 != len(portEnvVar) {
		port = portEnvVar
	}

	log.Printf("Starting webhook server on port %v ...", port)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Printf("Quit webhook server.")
}
