package glogger

import (
	"os"
	"time"
)

const (
	YYYYMMDD  = "02.01.2006"
	HHMMSS24h = "15:04:05"
	INF       = iota
	ERR
)

type Glogger struct {
	f *os.File
}

func (w *Glogger) Write(str string, mode int) (int, error) {
	ts := time.Now().UTC().Format(YYYYMMDD + " " + HHMMSS24h)
	wstr := "["

	switch mode {
	case INF:
		wstr += "INF"
	case ERR:
		wstr += "ERR"
	}

	wstr += "] " + ts + " | " + str + "\n" // [INF] YYYYMMDD HHMMSS24h | Текст лога

	return w.f.WriteString(wstr)
}

func Init(fileName string) (*Glogger, error) {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	return &Glogger{f}, nil
}

func (l *Glogger) Info(str string) {
	l.Write(str, INF)
}

func (l *Glogger) Error(err error) {
	l.Write(err.Error(), ERR)
}

func (l *Glogger) Close() error {
	return l.f.Close()
}

