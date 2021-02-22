package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"graceful-shutdown/internal/app"
)

func main() {
	timeWait := 15 * time.Second
	application := app.New()
	signChan := make(chan os.Signal, 1)
	go func() {
		if err := application.Start(); err != nil {
			log.Printf("%v", err.Error())
		}
	}()
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
	<-signChan
	log.Println("Shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), timeWait)
	defer cancel()
	if err := application.Stop(ctx); err == context.DeadlineExceeded {
		log.Print("Halted active connections")
	}
	close(signChan)
	log.Printf("Completed")
}
