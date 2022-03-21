package main

import (
	"github.com/nongfah/nmk/system"
	"os"
	"os/signal"
)

func main() {
	go func() {
		system.KeyboardMonitor(func(event system.Event) {
		})
	}()
	go func() {
		system.MouseMonitor(func(event system.Event) {
		})
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}
