package main

import (
	"log"
	"net/http"

	"github.com/walkersumida/go-api-server/controller"
	"github.com/walkersumida/go-api-server/ogen"
)

func main() {
	ctrl := controller.Controller{}

	srv, err := ogen.NewServer(ctrl)
	if err != nil {
			log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", srv); err != nil {
			log.Fatal(err)
	}
}
