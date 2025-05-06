package main

import (
	"log"

	"github.com/drizlye0/GopherSocial/internal/env"
	"github.com/drizlye0/GopherSocial/internal/store"
)

func main() {
	cfg := &config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: *cfg,
		store:  store,
	}

	mux := app.mount()
	err := app.run(mux)

	if err != nil {
		log.Fatal(err)
	}

}
