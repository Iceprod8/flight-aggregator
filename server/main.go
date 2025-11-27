package main

import (
	"aggregator/handler"
	"aggregator/repository"
	"aggregator/service"
	"aggregator/utils"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	port := config.SERVER_PORT

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.HealthHandler)

	repo1 := repository.NewJServer1Repository(config.JSERVER1_FILE_PATH) 
	repo2 := repository.NewJServer2Repository(config.JSERVER2_FILE_PATH)

	fs := service.NewFlightService(repo1, repo2)
	mux.HandleFunc("/flights", handler.FlightHandler(fs))

	// HTTP server configuration
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Server listening on :%s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	log.Println("Server stopped gracefully")
}
