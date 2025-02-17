package logger

import (
	"database/sql"
	"log"
)

const (
	Info = iota
	Error
)

var Writer Logger

type Logger struct {
	db *sql.DB
}

func SetGlobalLogger(db *sql.DB) {
	Writer = *NewLogger(db)
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
