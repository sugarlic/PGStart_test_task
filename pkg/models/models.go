package models

import "errors"

var ErrNoRecord = errors.New("models: no current command")

type Command struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Exec_res string `json:"exec_res"`
}
