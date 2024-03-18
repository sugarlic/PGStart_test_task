package models

import "errors"

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Command struct {
	ID      int
	Title   string
	Content string
}
