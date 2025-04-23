package main

import (
	"github.com/gorilla/mux"
	"github.com/lahiruudayakumara/bet_settlement_engine/api/routes"
	"github.com/lahiruudayakumara/bet_settlement_engine/config"
	"github.com/lahiruudayakumara/bet_settlement_engine/internal/store"
	"github.com/lahiruudayakumara/bet_settlement_engine/scripts"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// WelcomeHandler handles the root route and provides a welcome message
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Bet Settlement Engine API"))
}

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize in-memory store
	store := store.NewInMemoryStore()
	// Seed initial data
	scripts.SeedData(store)

	// Create a new router using gorilla/mux
	r := mux.NewRouter()

	// Define the Welcome route
	r.HandleFunc("/", WelcomeHandler).Methods("GET")

	// Register other routes
	routes.RegisterBetRoutes(r)
	routes.RegisterUserRoutes(r)
	routes.RegisterEventRoutes(r)
	routes.RegisterBetTransactionRoutes(r)
	routes.RegisterBetResultRoutes(r)

	// Configure server with graceful shutdown support
	port := "8089"
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	srv := &http.Server{
		Handler:      r,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Channel to handle graceful shutdown signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Run the server in a goroutine
	go func() {
		log.Printf("Server starting on port %v...", config.AppConfig.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %v: %v", config.AppConfig.Server.Port, err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	<-stop

	// Gracefully shutdown the server
	log.Println("Shutting down server...")
	if err := srv.Close(); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}
