package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofor-little/env"
)

func main() {
	fmt.Println("Initializing server")

	if err := env.Load("main.env"); err != nil {
		log.Fatalln(err)
	}

	srv := ServerSetup()

	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println("Error when trying to serve:", err, srv.Addr)
		}
	}()

	log.Print("Server is up and ready @ ", srv.Addr)

	<-stopChannel

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalln("Error when shutting down:", err)
	}

	log.Println("Server has been shut down")
}
