package main

import (
	"time"

	"github.com/drizlye0/GopherSocial/internal/db"
	"github.com/drizlye0/GopherSocial/internal/env"
	"github.com/drizlye0/GopherSocial/internal/mailer"
	"github.com/drizlye0/GopherSocial/internal/store"
	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			GohperSocial API
//	@description	This is a social network API for Gophers!
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description

func main() {
	cfg := &config{
		addr:        env.GetString("ADDR", ":8080"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:3000/"),
		frontendURL: env.GetString("FRONTEND_URL", "localhost:4000/"),
		db: dbConfig{
			addr:        env.GetString("DB_ADDR", "postgres://user:adminpass@localhost/social?sslmode=disabled"),
			maxOpenConn: env.GetInt("DB_MAX_OPEN_CONN", 30),
			maxIdleConn: env.GetInt("DB_MAX_IDLE_CONN", 30),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp:       time.Hour * 24 * 3, // 3 days
			fromEmail: env.GetString("FROM_EMAIL", ""),
			sengrid: sendGridConfig{
				apiKey: env.GetString("SENGRID_API_KEY", ""),
			},
		},
		auth: authConfig{
			basic: basicConfig{
				user: env.GetString("AUTH_BASIC_USER", "admin"),
				pass: env.GetString("AUTH_BASIC_PASS", "admin"),
			},
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConn,
		cfg.db.maxIdleConn,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("db connection pool established")

	store := store.NewStorage(db)

	mailer := mailer.NewSendgrid(cfg.mail.sengrid.apiKey, cfg.mail.fromEmail)

	app := &application{
		config: *cfg,
		store:  store,
		logger: logger,
		mailer: mailer,
	}

	mux := app.mount()
	err = app.run(mux)

	if err != nil {
		logger.Fatal(err)
	}

}
