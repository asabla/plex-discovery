package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/asabla/plex-discovery/internal/plex"
	"github.com/asabla/plex-discovery/internal/server"
	"github.com/asabla/plex-discovery/internal/ui/pages"
)

func main() {
	cfg := server.Config{
		Address:         envOrDefault("SERVER_ADDRESS", ":8080"),
		ShutdownTimeout: 10 * time.Second,
	}

	plexClient, err := plex.NewClient(envOrDefault("PLEX_BASE_URL", "https://plex.tv"), nil)
	if err != nil {
		log.Fatalf("failed to create Plex client: %v", err)
	}

	mux := server.NewRouter(pages.Home(), plexClient)
	httpServer := server.NewHTTPServer(cfg, mux)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("server listening on %s", cfg.Address)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
		return
	}
	log.Println("server stopped")
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
