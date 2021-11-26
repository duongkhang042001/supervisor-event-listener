package main

import (
	"supervisor-event-listener/listener"
)

func main() {
	for {
		listener.Start()
	}
}
