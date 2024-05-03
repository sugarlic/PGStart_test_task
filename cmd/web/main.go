package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/test/pkg/models/postgre"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	commands postgre.CommandService
}

func main() {
	addr := flag.String("addr", ":8080", "Сетевой адрес веб-сервера")
	dsn := flag.String("dsn", "host=localhost port=5432 user=postgres password=1 dbname=postgres sslmode=disable", "Название PostgreSQL источника данных")
	flag.Parse()

	// infoLogFile, err := os.OpenFile("./log/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatalf("error opening info log file: %v", err)
	// }
	// defer infoLogFile.Close()

	// errorLogFile, err := os.OpenFile("./log/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatalf("error opening error log file: %v", err)
	// }
	// defer errorLogFile.Close()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dbpool, err := openDBpool(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer dbpool.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		commands: &postgre.CommandModel{DB: dbpool},
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

func openDBpool(dsn string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	if err = dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return dbpool, nil
}
