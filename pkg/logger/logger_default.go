package logger

import (
	"database/sql"
	"log"
)

const (
	Info = iota
	Error
)

type Logger struct {
	db *sql.DB
}

func NewLogger(db *sql.DB) *Logger {
	return &Logger{db: db}
}

func (l *Logger) Info(message string) {
	l.registerLog(message, Info)
}

func (l *Logger) Error(message string) {
	l.registerLog(message, Error)
}

func (l *Logger) registerLog(message string, level int) {
	q := "INSERT INTO logs (message, level) VALUES (?, ?)"
	if _, err := l.db.Exec(q, message, level); err != nil {
		panic(err)
	}

	log.Println(message)
}
