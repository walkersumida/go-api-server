package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/walkersumida/go-api-server/database"
	"github.com/walkersumida/go-api-server/ent"
	"github.com/walkersumida/go-api-server/handler"
	"github.com/walkersumida/go-api-server/ogen"
	"github.com/walkersumida/go-api-server/repository"
)

func main() {
	rootCmd := &cobra.Command{
		Use:  "server",
		Args: cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			return run()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	url := database.BuildDBURL(database.Configs{})
	entCli, err := ent.Open(dialect.Postgres, url)
	if err != nil {
		return errors.Wrap(err, "failed opening connection to postgres")
	}
	defer entCli.Close()

	repo := repository.NewRepository(entCli)

	h := handler.NewHandler(repo)

	srv, err := ogen.NewServer(h)
	if err != nil {
		return errors.Wrap(err, "failed to run server")
	}

	const timeout = 60 * time.Second
	httpServer := http.Server{
		Addr:         ":8080",
		Handler:      srv,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	idleConnsClosed := make(chan struct{})
	go func() {
		<-ctx.Done()

		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Println("Server started")
	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
	log.Println("Server stopped")

	return nil
}
