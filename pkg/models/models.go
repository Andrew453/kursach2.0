package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Comment struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
