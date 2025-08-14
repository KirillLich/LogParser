package filter

import (
	"strings"
	"time"

	"github.com/KirillLich/logparser/pkg/parse"
)

func Filter(records []parse.LogRecord, level string, start time.Time, end time.Time, substr string) []parse.LogRecord {
	resultRecords := records
	if level != "all" {
		resultRecords = ByLevel(resultRecords, level)
	}
	if (start != time.Time{}) && (end != time.Time{}) {
		resultRecords = ByTime(resultRecords, start, end)
	}
	if substr != "" {
		resultRecords = ByContains(resultRecords, substr)
	}
	return resultRecords
}

func ByLevel(records []parse.LogRecord, level string) []parse.LogRecord {
	var resultRecords []parse.LogRecord
	level = strings.ToUpper(level)
	for _, record := range records {
		if record.Level == level {
			resultRecords = append(resultRecords, record)
		}
	}
	return resultRecords
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func ByTime(records []parse.LogRecord, start time.Time, end time.Time) []parse.LogRecord {
	var resultRecords []parse.LogRecord
	for _, record := range records {
		if inTimeSpan(start, end, record.TimeStamp) {
			resultRecords = append(resultRecords, record)
		}
	}
	return resultRecords
}

func ByContains(records []parse.LogRecord, substr string) []parse.LogRecord {
	var resultRecords []parse.LogRecord
	for _, record := range records {
		if strings.Contains(record.Message, substr) {
			resultRecords = append(resultRecords, record)
		}
	}
	return resultRecords
}
