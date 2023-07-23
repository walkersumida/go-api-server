package main

import (
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
	"github.com/walkersumida/go-api-server/database"
	"github.com/walkersumida/go-api-server/ent"
	"github.com/walkersumida/go-api-server/handler"
	"github.com/walkersumida/go-api-server/ogen"
	"github.com/walkersumida/go-api-server/repository"
)

func main() {
	url := database.BuildDBURL(database.Configs{})
	entCli, err := ent.Open(dialect.Postgres, url)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer entCli.Close()

	repo := repository.NewRepository(entCli)

	h := handler.NewHandler(repo)

	srv, err := ogen.NewServer(h)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
