package handler

import (
	"log"
	"net/http"
	"os/exec"
)

const doneCommand = `/usr/local/bin/t30-on`

//Done A handler for octorpint server print done event.
func Done(w http.ResponseWriter, r *http.Request) {
	delays, ok := r.URL.Query()["delay"]

	delayInMinutes := "0"
	if !ok || len(delays[0]) < 1 {
		log.Println("No delay!")
	} else {
		delayInMinutes = delays[0]
	}

	cmd := exec.Command("/usr/bin/at", "now", "+", delayInMinutes, "minutes", "-f", doneCommand)
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to run command with %s\n", err)
	} else {
		log.Printf("Succeeded in executing command\n")
	}
}
