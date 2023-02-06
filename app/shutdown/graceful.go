package shutdown

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Gracefully() {
	// Setting up signal capturing
	quit := make(chan os.Signal, 1)
	defer close(quit)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	// Waiting for SIGINT (pkill -2)
	<-quit
	log.Println("Graceful Shutdown start")
}
