package models

import "errors"

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Command struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
