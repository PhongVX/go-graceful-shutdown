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
	//1. Cho ứng dụng của chúng ta chạy background trong 1 Goroutine
	go func() {
		if err := application.Start(); err != nil {
			log.Printf("%v", err.Error())
		}
	}()
	//2. Thiết lập một channel để lắng nghe tín hiệu dừng từ hệ điều hành, 
	//   ở đây chúng ta lưu ý 2 tín hiệu (signal) là SIGINT và SIGTERM
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
	<-signChan
	log.Println("Shutting down")
	//3. Thiết lập một khoản thời gian (Timeout) để dừng hoàn toàn ứng dụng và đóng tất cả kết nối.
	ctx, cancel := context.WithTimeout(context.Background(), timeWait)
	defer func() {
		log.Println("Close another connection")
		cancel()
	}()
	log.Println("Stop http server")
	if err := application.Stop(ctx); err == context.DeadlineExceeded {
		log.Print("Halted active connections")
	}
	close(signChan)
	log.Printf("Completed")
}
