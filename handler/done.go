package handler

import (
	"log"
	"net/http"
	"os/exec"
)

// change me!
const doneCommand = `/usr/local/bin/s2-off`

//Done A handler for octorpint server print done event.
func Done(w http.ResponseWriter, r *http.Request) {
	delays, ok := r.URL.Query()["delay"]
	delayInMinutes := "0"
	if ok && 0 != len(delays[0]) {
		delayInMinutes = delays[0]
	}
	log.Printf("Delay for %v minutes\n", delayInMinutes)
	cmd := exec.Command("/usr/bin/at", "now", "+", delayInMinutes, "minutes", "-f", doneCommand)
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to run command with %s\n", err)
	} else {
		log.Printf("Succeeded in executing command\n")
	}
}
