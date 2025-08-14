package parse

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/KirillLich/logparser/internal/config"
)

// Maybe I can make separate type of logRaw if I do not use it in LogRecord struct
// this may affect realisation of bindByFields func
type LogRecord struct {
	logRaw    map[string]any
	Level     string
	Message   string
	TimeStamp time.Time
}

// Option for log file < your RAM capacity
func ReadLogFile(path string, cfg config.Config) ([]LogRecord, error) {
	logFile, err := os.ReadFile(path)
	if err != nil {
		return []LogRecord{}, err
	}

	RawLogs := strings.Split(string(logFile), "\n")
	var ResultLogs []LogRecord

	for _, RawLog := range RawLogs {
		if RawLog == "" {
			continue
		}

		var raw map[string]any
		err := json.Unmarshal([]byte(RawLog), &raw)

		//change from throwing error to just logging it and continue

		//its necessary to wrap error then sending it
		if err != nil {
			return ResultLogs, fmt.Errorf("unmarshal error: %w", err)
		}

		record := LogRecord{logRaw: raw}
		if err := bindByFields(&record, cfg.Fields); err != nil {
			return ResultLogs, fmt.Errorf("binding fields error: %w", err)
		}

		ResultLogs = append(ResultLogs, record)
	}

	return ResultLogs, nil
}

// can try making array of errors and throw it all at once
func bindByFields(record *LogRecord, fields config.Fields) error {
	LogRaw := record.logRaw
	//TODO: add error catching while type assertion
	record.Level = LogRaw[fields.Level].(string)
	record.Message = LogRaw[fields.Message].(string)
	TimeRaw := LogRaw[fields.Time].(string)
	var err error
	record.TimeStamp, err = time.Parse(fields.TimeLayout, TimeRaw)
	return err
}
