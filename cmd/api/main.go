package main

import (
	"log"

	"github.com/drizlye0/GopherSocial/internal/env"
)

func main() {
	cfg := &config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: *cfg,
	}

	mux := app.mount()
	err := app.run(mux)

	if err != nil {
		log.Fatal(err)
	}

}
