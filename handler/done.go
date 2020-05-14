package handler

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

// change me!
const (
	defaultDoneCommand    = `/usr/local/bin/s2-off`
	doneCommandEnvVarName = "DONE_COMMAND"
)

//Done A handler for octorpint server print done event.
func Done(w http.ResponseWriter, r *http.Request) {
	delays, ok := r.URL.Query()["delay"]
	delayInMinutes := "0"
	if ok && 0 != len(delays[0]) {
		delayInMinutes = delays[0]
	}

	doneCommandEnv := os.Getenv(doneCommandEnvVarName)
	doneCommand := defaultDoneCommand
	if 0 != len(doneCommandEnv) {
		doneCommand = doneCommandEnv
	}
	log.Printf("Execute command %v with delay %v minutes.\n", doneCommand, delayInMinutes)
	cmd := exec.Command("/usr/bin/at", "now", "+", delayInMinutes, "minutes", "-f", doneCommand)
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to run command with error: %s\n", err)
	} else {
		log.Printf("Succeeded in executing command.\n")
	}
}
