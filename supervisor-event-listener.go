package main

import (
	"github.com/lwldcr/supervisor-event-listener/listener"
)

func main() {
	for {
		listener.Start()
	}
}
