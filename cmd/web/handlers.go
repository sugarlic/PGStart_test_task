package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/test/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.commands.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, s)
}

func (app *application) getCommand(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.commands.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.infoLog.Println(s)

	app.render(w, s)
}

func (app *application) createCommand(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		app.serverError(w, err)
		return
	}

	var data models.Command
	err = json.Unmarshal(body, &data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = app.commands.Insert(data.Title, data.Content)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write([]byte("Succes"))
}
