package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"greenlight.alexedwards.net/internal/data"
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
}

type application struct {
	config conf
	logger *log.Logger
	models data.Models
}

func main() {

	var config conf

	flag.IntVar(&config.port, "port", 4000, "API server port")
	flag.StringVar(&config.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&config.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "postgres dsn")

	flag.IntVar(&config.db.maxOpenConn, "db-max-open-conns", 25, "postgres max open conn")
	flag.IntVar(&config.db.maxIdleConn, "db-max-idle-conns", 25, "postgres max idle conn")
	flag.StringVar(&config.db.maxIdleTime, "db-max-idle-time", "15m", "postgres max idle time")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logger.Printf("database connection pool established")

	app := &application{
		config: config,
		logger: logger,
		models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.port),
		Handler:      app.router(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on port %d ", config.env, config.port)
	err = srv.ListenAndServe()
	logger.Fatal(err)
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
