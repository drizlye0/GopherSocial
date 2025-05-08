package main

import (
	"fmt"
	"log"

	"github.com/drizlye0/GopherSocial/internal/db"
	"github.com/drizlye0/GopherSocial/internal/env"
	"github.com/drizlye0/GopherSocial/internal/store"
)

const version = "0.0.1"

func main() {
	cfg := &config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:        env.GetString("DB_ADDR", "postgres://user:adminpass@localhost/social?sslmode=disabled"),
			maxOpenConn: env.GetInt("DB_MAX_OPEN_CONN", 30),
			maxIdleConn: env.GetInt("DB_MAX_IDLE_CONN", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
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
	fmt.Println("db connection pool established")

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
