package main
import (
	"fmt"
	"time"
	"os"
	"github.com/stianeikeland/go-rpio/v4"
)

var (
	pin = rpio.Pin(10)
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	defer rpio.Close()

	pin.Output()

	for {
		pin.Toggle()
		time.Sleep(time.Second)
	}
}

