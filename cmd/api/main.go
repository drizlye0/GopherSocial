package main

import (
	"log"

	"github.com/drizlye0/GopherSocial/internal/db"
	"github.com/drizlye0/GopherSocial/internal/env"
	"github.com/drizlye0/GopherSocial/internal/store"
)

func main() {
	cfg := &config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:        env.GetString("DB_ADDR", "postgres://user:adminpass@localhost/social?sslmode=diabled"),
			maxOpenConn: env.GetInt("DB_MAX_OPEN_CONN", 30),
			maxIdleConn: env.GetInt("DB_MAX_IDLE_CONN", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConn,
		cfg.db.maxIdleConn,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	store := store.NewStorage(db)

	app := &application{
		config: *cfg,
		store:  store,
	}

	mux := app.mount()
	err = app.run(mux)

	if err != nil {
		log.Fatal(err)
	}

}
