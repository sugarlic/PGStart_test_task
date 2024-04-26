package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/test/pkg/models/postgre"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	commands postgre.CommandService
}

func main() {
	addr := flag.String("addr", ":8080", "Сетевой адрес веб-сервера")
	dsn := flag.String("dsn", "host=db port=5432 user=postgres password=1 dbname=postgres sslmode=disable", "Название PostgreSQL источника данных")
	flag.Parse()

	infoLogFile, err := os.OpenFile("./log/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening info log file: %v", err)
	}
	defer infoLogFile.Close()

	errorLogFile, err := os.OpenFile("./log/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening error log file: %v", err)
	}
	defer errorLogFile.Close()

	infoLog := log.New(infoLogFile, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(errorLogFile, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		commands: &postgre.CommandModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Запуск сервера на %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
