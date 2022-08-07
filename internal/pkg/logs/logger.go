package logs

import (
	"fmt"
	domLogger "github-user-service/internal/domain/adaptors/logger"
	"log"
)

const (
	FATAL string = `FATAL`
	ERROR string = `ERROR`
	WARN  string = `WARN`
	INFO  string = `INFO`
	DEBUG string = `DEBUG`
	TRACE string = `TRACE`
)

// LevelMap in-memory map was kept to track the levels of the logs
// string type was used to increase the visibility of the log type
var LevelMap = map[string]int{
	FATAL: 6,
	ERROR: 5,
	WARN:  4,
	INFO:  3,
	DEBUG: 2,
	TRACE: 1,
}

type logger struct {
	Level string
}

func NewLogger(level string) (domLogger.Logger, error) {
	_, ok := LevelMap[level]
	if !ok {
		return nil, fmt.Errorf("invalid log level received [%s]", level)
	}
	return &logger{
		Level: level,
	}, nil
}

func (l logger) Fatal(message string) {
	log.Fatalln(message)
}

func (l logger) Error(message string) {
	if LevelMap[l.Level] < LevelMap[ERROR] {
		log.Println("[ERROR] ", message)
	}
}

func (l logger) Warn(message string) {
	if LevelMap[l.Level] < LevelMap[WARN] {
		log.Println("[WARN] ", message)
	}
}

func (l logger) Debug(message string) {
	if LevelMap[l.Level] < LevelMap[DEBUG] {
		log.Println("[DEBUG] ", message)
	}
}

func (l logger) Info(message string) {
	if LevelMap[l.Level] < LevelMap[INFO] {
		log.Println("[INFO] ", message)
	}
}

func (l logger) Trace(message string) {
	if LevelMap[l.Level] < LevelMap[TRACE] {
		log.Println("[TRACE] ", message)
	}
}
