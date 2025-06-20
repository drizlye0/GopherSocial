package main

import (
	"expvar"
	"runtime"
	"time"

	"github.com/drizlye0/GopherSocial/internal/auth"
	"github.com/drizlye0/GopherSocial/internal/db"
	"github.com/drizlye0/GopherSocial/internal/env"
	"github.com/drizlye0/GopherSocial/internal/mailer"
	"github.com/drizlye0/GopherSocial/internal/ratelimiter"
	"github.com/drizlye0/GopherSocial/internal/store"
	"github.com/drizlye0/GopherSocial/internal/store/cache"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

const version = "1.1.0"

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
		redisCfg: redisConfig{
			addr:    env.GetString("REDIS_ADDR", "localhost:6379"),
			pw:      env.GetString("REDIS_PW", ""),
			db:      env.GetInt("REDIS_DB", 0),
			enabled: env.GetBool("REDIS_ENABLED", true),
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
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "example"),
				exp:    time.Hour * 24 * 3, // 3 Days
				iss:    "gophersocial",
			},
		},
		rateLimiter: ratelimiter.Config{
			RequestPerTimeFrame: env.GetInt("RATELIMITER_REQUEST_COUNT", 20),
			TimeFrame:           time.Second * 5,
			Enabled:             env.GetBool("RATELIMITER_ENABLED", true),
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

	var rdb *redis.Client
	if cfg.redisCfg.enabled {
		rdb = cache.NewRedisClient(cfg.redisCfg.addr, cfg.redisCfg.pw, cfg.redisCfg.db)
		logger.Info("redis cahe connection established")
	}

	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.rateLimiter.RequestPerTimeFrame,
		cfg.rateLimiter.TimeFrame,
	)

	store := store.NewStorage(db)
	cacheStorage := cache.NewRedisStorage(rdb)

	mailer := mailer.NewSendgrid(cfg.mail.sengrid.apiKey, cfg.mail.fromEmail)

	jwtAuthenticator := auth.NewJWTAuthenticator(
		cfg.auth.token.secret,
		cfg.auth.token.iss,
		cfg.auth.token.iss,
	)

	app := &application{
		config:        *cfg,
		cacheStorage:  cacheStorage,
		store:         store,
		logger:        logger,
		mailer:        mailer,
		authenticator: jwtAuthenticator,
		rateLimiter:   rateLimiter,
	}

	// metrics collected
	expvar.NewString("version").Set(version)
	expvar.Publish("database", expvar.Func(func() any {
		return db.Stats()
	}))
	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	mux := app.mount()
	err = app.run(mux)
	if err != nil {
		logger.Fatal(err)
	}
}
