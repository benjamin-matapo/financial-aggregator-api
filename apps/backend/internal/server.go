package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"financial-aggregator-api/backend/handlers"
	"financial-aggregator-api/backend/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Server represents the HTTP server
type Server struct {
	router *chi.Mux
	server *http.Server
}

// NewServer creates a new Server instance
func NewServer() *Server {
	// Initialize services
	accountService := services.NewAccountService()
	transactionService := services.NewTransactionService()

	// Initialize handlers
	accountHandler := handlers.NewAccountHandler(accountService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Create router
	router := chi.NewRouter()

	// Add middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Timeout(60 * time.Second))

	// CORS configuration
	// CORS: allow all origins for simplicity on Render/preview
	// TODO: In production, restrict to your frontend origin(s)
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})
	router.Use(corsConfig.Handler)

	// Health check endpoint
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[health] %s %s", r.Method, r.URL.String())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":    "healthy",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// API routes
	router.Route("/api", func(r chi.Router) {
		// Account routes
		r.Route("/accounts", func(r chi.Router) {
			r.Get("/", accountHandler.GetAccounts)
			r.Get("/{id}", accountHandler.GetAccountByID)
			r.Post("/{id}/refresh", accountHandler.RefreshAccount)
			r.Get("/{id}/transactions", transactionHandler.GetTransactionsByAccount)
		})

		// Transaction routes
		r.Route("/transactions", func(r chi.Router) {
			r.Get("/", transactionHandler.GetTransactions)
			r.Get("/{id}", transactionHandler.GetTransactionByID)
		})
	})

	return &Server{
		router: router,
	}
}

// Start starts the HTTP server
func (s *Server) Start(port string) error {
	if port == "" {
		port = "8080"
	}

	s.server = &http.Server{
		Addr:    ":" + port,
		Handler: s.router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server is shutting down...")

	// Give outstanding requests 30 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
	return nil
}

// GetRouter returns the router for testing purposes
func (s *Server) GetRouter() *chi.Mux {
	return s.router
}
