package safety

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func GracefullyStopBootstrap(bootstrap func()) {
	bootstrap()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down local listener...")
	log.Println("Local listener stopped.")
}