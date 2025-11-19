package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/notabot/backend/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Oracle service starting with config: %+v", cfg)

	// TODO: Initialize oracle service
	// - Blockchain connection (Base/Ethereum)
	// - Contract interaction
	// - Event listeners
	// - Transaction submission

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down oracle service...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// TODO: Graceful shutdown of oracle services
	_ = ctx

	log.Println("Oracle service exited")
}


