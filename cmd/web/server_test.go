package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/test/pkg/models/postgre"
)

var app = &application{
	errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
}

var dsn string = "host=localhost port=5432 user=postgres password=1 dbname=postgres sslmode=disable"

func TestHomeHandler(t *testing.T) {
	// создание экземпляра приложения
	db, err := openDB(dsn)
	if err != nil {
		app.errorLog.Fatal(err)
	}

	defer db.Close()

	app.commands = &postgre.CommandModel{DB: db}

	// тестировка обработчика /
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.home)

	handler.ServeHTTP(rr, req)

	// Проверка кода ответа
	if rr.Code != http.StatusOK {
		t.Errorf("Ожидаемый код ответа %d, полученный %d", http.StatusOK, rr.Code)
	}
}
