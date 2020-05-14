package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/rickyzhang82/octoprint-webhook-server/handler"
)

const (
	defaultPort    = "12000"
	portEnvVarName = "WEBHOOK_SEREVER_PORT"
)

func appCleanup() {
	log.Printf("Quit webhook server.")
	os.Exit(0)
}

func main() {
	http.HandleFunc("/done", handler.Done)
	portEnvVar := os.Getenv(portEnvVarName)
	port := defaultPort
	if 0 != len(portEnvVar) {
		port = portEnvVar
	}
	// setup signal catching
	sigs := make(chan os.Signal, 1)

	// catch all signals since not explicitly listing
	signal.Notify(sigs)
	//signal.Notify(sigs,syscall.SIGQUIT)

	// method invoked upon seeing signal
	go func() {
		s := <-sigs
		log.Printf("RECEIVED SIGNAL: %s", s)
		appCleanup()
	}()

	log.Printf("Starting webhook server on port %v ...", port)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		appCleanup()
	}
}
