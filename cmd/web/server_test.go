package main

import (
	"encoding/json"
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

func TestHomeNotFound(t *testing.T) {
	app := newTestApplication(t)

	req, _ := http.NewRequest("GET", "/a", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getCommand)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, rr.Code)
	}
}

func TestHomeList(t *testing.T) {
	mockCommands := &mocks.MockCommands{
		Commands:   make([]*models.Command, 0),
		ExpectCall: make(map[string]int),
	}

	tests := []*models.Command{{
		ID:       1,
		Title:    "test1",
		Content:  "test1",
		Exec_res: "test1",
	}, {
		ID:       2,
		Title:    "aboba",
		Content:  "aboba_content",
		Exec_res: "aboba_res",
	}, {
		ID:       3,
		Title:    "school21",
		Content:  "school21_content",
		Exec_res: "school21_res",
	}}

	mockCommands.Commands = append(mockCommands.Commands, tests...)

	app := &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
		commands: mockCommands,
	}

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.home)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rr.Code)
	}

	rr.Body.Bytes()
	var res []*models.Command
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		t.Error("Error with response reading")
	}

	if len(res) != len(tests) {
		t.Error("Error with response len")
	}

	for i := 0; i < len(res); i++ {
		if tests[i].ID != res[i].ID {
			t.Errorf("Expected ID %d, got %d at %d iteration", tests[i].ID, res[i].ID, i)
		}
		if tests[i].Title != res[i].Title {
			t.Errorf("Expected ID %s, got %s at %d iteration", tests[i].Title, res[i].Title, i)
		}
		if tests[i].Content != res[i].Content {
			t.Errorf("Expected ID %s, got %s at %d iteration", tests[i].Content, res[i].Content, i)
		}
		if tests[i].Exec_res != res[i].Exec_res {
			t.Errorf("Expected ID %s, got %s at %d iteration", tests[i].Exec_res, res[i].Exec_res, i)
		}
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
