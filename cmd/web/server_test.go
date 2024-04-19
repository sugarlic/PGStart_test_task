package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/test/mocks"
	_ "github.com/test/mocks"
	"github.com/test/pkg/models"
)

func newTestApplication(t *testing.T) *application {
	// Создание моков для команд
	mockCommands := &mocks.MockCommands{
		Commands:   make([]*models.Command, 0),
		ExpectCall: make(map[string]int),
	}

	// Создание тестовой структуры application
	return &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
		commands: mockCommands,
	}
}

func TestGetCommandNotFound(t *testing.T) {
	app := newTestApplication(t)

	req, _ := http.NewRequest("GET", "/command?id=1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getCommand)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, rr.Code)
	}
}
