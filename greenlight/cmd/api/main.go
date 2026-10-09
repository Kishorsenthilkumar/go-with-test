package main

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"time"

	_ "github.com/lib/pq"
	"greenlight.alexedwards.net/internal/data"
	"greenlight.alexedwards.net/internal/jsonlog"
	"greenlight.alexedwards.net/internal/mailer"
)

const version = "1.0.0"

type conf struct {
	port int
	env  string
	db   struct {
		dsn         string
		maxOpenConn int
		maxIdleConn int
		maxIdleTime string
	}
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

type application struct {
	config conf
	logger jsonlog.Logger
	models data.Models
	mailer mailer.Mailer
}

func main() {

	var config conf

	flag.IntVar(&config.port, "port", 4000, "API server port")
	flag.StringVar(&config.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&config.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "postgres dsn")

	flag.IntVar(&config.db.maxOpenConn, "db-max-open-conns", 25, "postgres max open conn")
	flag.IntVar(&config.db.maxIdleConn, "db-max-idle-conns", 25, "postgres max idle conn")
	flag.StringVar(&config.db.maxIdleTime, "db-max-idle-time", "15m", "postgres max idle time")

	flag.Float64Var(&config.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&config.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&config.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.StringVar(&config.smtp.host, "smtp-host", "live.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&config.smtp.port, "smtp-port", 587, "SMTP port")
	flag.StringVar(&config.smtp.username, "smtp-username", "smtp@mailtrap.io", "SMTP username")
	flag.StringVar(&config.smtp.password, "smtp-password", "e4d5e0a0ea74582b924089e8ee0aaf0d", "SMTP password")
	flag.StringVar(&config.smtp.sender, "smtp-sender", "hello@demomailtrap.co", "SMTP sender")

	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(config)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()

	logger.PrintInfo("database connection pool established", nil)

	app := &application{
		config: config,
		logger: *logger,
		models: data.NewModels(db),
		mailer: mailer.New(config.smtp.host, config.smtp.port, config.smtp.username, config.smtp.password, config.smtp.sender),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}

func openDB(config conf) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.db.maxOpenConn)
	db.SetMaxIdleConns(config.db.maxIdleConn)

	duration, err := time.ParseDuration(config.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
