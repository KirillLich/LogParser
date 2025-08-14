package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/KirillLich/logparser/internal/config"
	customflags "github.com/KirillLich/logparser/internal/customFlags"
	"github.com/KirillLich/logparser/internal/filter"
	"github.com/KirillLich/logparser/pkg/parse"
)

func main() {
	//TODO: move path into env variables
	cfg := config.MustLoad("../configs/local.yaml")

	fmt.Println("Hello from parser. Here is your config:", cfg)

	//TODO: make different logs levels
	log := setupJSONLogger("../logs/app.json", &slog.HandlerOptions{Level: slog.LevelDebug})

	log.Debug("Hello from logger")

	//TODO: make custom log file with flags
	example, err := parse.ReadLogFile("../example/app.json", cfg)
	if err != nil {
		log.Error("error while parsing log file.", slog.String("Error", err.Error()))
	}
	log.Debug("Here is your last log data", slog.String("", example[len(example)-1].Message))

	//countFlag := flag.Bool("count", false, "flag for counting mode")
	levelFlag := flag.String("level", "all", "flag for filter.ByLevel")
	containsFlag := flag.String("contains", "", "flag for filter.ByContains")

	before, after := &time.Time{}, &time.Time{}
	//TODO: add usage string
	flag.Var(&customflags.TimeValue{before, cfg.TimeLayout}, "before", "")
	flag.Var(&customflags.TimeValue{after, cfg.TimeLayout}, "after", "")

	flag.Parse()

	result := filter.Filter(example, *levelFlag, *before, *after, *containsFlag)

	for _, r := range result {
		fmt.Printf("[%s] %s — %s\n", r.TimeStamp.Format(cfg.TimeLayout), r.Level, r.Message)
	}
}

func setupJSONLogger(logFilePath string, opts *slog.HandlerOptions) *slog.Logger {
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	handler := slog.NewJSONHandler(logFile, opts)

	return slog.New(handler)
}

func printExample(logArray []parse.LogRecord) {
	for _, log := range logArray {
		fmt.Println(log)
	}
}
