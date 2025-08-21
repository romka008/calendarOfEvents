package cmd

import (
	"sync"
	"time"
)

type Logger struct {
	mu   sync.Mutex
	list []Log
}

type Log struct {
	At    string
	Value string
}

func (l *Logger) AddLog(str string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list = append(l.list, Log{At: time.Now().String(), Value: str})
}

func (l *Logger) GetLog() []Log {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.list
}
