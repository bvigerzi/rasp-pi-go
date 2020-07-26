package main
import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"github.com/stianeikeland/go-rpio/v4"
)

type state struct {
	On bool
}

var (
	pin = rpio.Pin(10)
	simulate = false
	on = false
)

func toggle() {
	if simulate {
		fmt.Println("Toggled the pin")
	} else {
		pin.Toggle()
	}
	on = !on
}

func unlock() {
	if simulate {
		fmt.Println("Unlocking")
	} else {
		pin.High()
		time.Sleep(time.Second)
		pin.Low()
	}
}

func init() {
	fmt.Println("Init the pin")
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		fmt.Println("Failed to init GPIO pin, simulating instead")
		simulate = true
	} else {
		fmt.Println("Finished pin init")
		pin.Output()
		pin.Low()
	}
}

func toggleHandler(responseWriter http.ResponseWriter, request *http.Request) {
	toggle()
	responseWriter.WriteHeader(http.StatusOK)
	state := state{on}
	json.NewEncoder(responseWriter).Encode(state)
}

func unlockHandler(responseWriter http.ResponseWriter, request *http.Request) {
	unlock()
	responseWriter.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/toggle", toggleHandler)

	http.HandleFunc("/unlock", unlockHandler)

	http.ListenAndServe(":8080", nil)
}
