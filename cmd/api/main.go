package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Bruheem/Portail_de_Reservation/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	jwt  struct {
		secret    string
		issuer    string
		port      int
		tokenLife time.Duration
	}
}

type application struct {
	config config
	logger *log.Logger

	library  *models.LibraryModel
	document *models.DocumentModel
	user     *models.UserModel
	lending  *models.LendingModel
}

func main() {
	var cfg config

	dsn := flag.String("dsn", "ibrahim:ibrahim@/PDRE?parseTime=true", "MySQL date source name")
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Server  environment")
	flag.Parse()

	logger := log.New(os.Stdout, "INFO:\t", log.Ltime)

	db, err := openDB(*dsn)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("starting database server")
	defer db.Close()

	app := &application{
		config:   cfg,
		logger:   logger,
		library:  &models.LibraryModel{DB: db},
		user:     &models.UserModel{DB: db},
		document: &models.DocumentModel{DB: db},
		lending:  &models.LendingModel{DB: db},
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on port %d", cfg.env, cfg.port)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
