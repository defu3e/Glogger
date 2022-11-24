package glogger

import (
	"log"
	"os"
)

type Glogger struct {
	f *os.File
}

func Init(fileName string) (*Glogger, error) {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	log.SetOutput(f)

	return &Glogger{f}, nil
}

func (l *Glogger) Info(str string) {
	log.Println(str)
}

func (l *Glogger) Error(err error) {
	log.Println(err)
}

func (l *Glogger) Close() error {
	return l.f.Close()
}

