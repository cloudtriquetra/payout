package utils

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type EffortEntry struct {
	EmployeeName   string
	StartTimeStamp time.Time
	EndTimeStamp   time.Time
	Duration       time.Duration
	Description    string
	EffortType     string
	CreatedAt      time.Time
	CreatedBy      string
}

func NewEffortEntry(empName string, startTimeStamp time.Time, endTimeStamp time.Time, duration time.Duration) (*EffortEntry, error) {
	if empName == "" {
		return nil, errors.New("employee Name is required")
	}
	return &EffortEntry{
		EmployeeName:   empName,
		StartTimeStamp: startTimeStamp,
		EndTimeStamp:   endTimeStamp,
		Duration:       duration,
		CreatedAt:      time.Now(),
	}, nil

}

func (effort EffortEntry) RecordEffortToFile(recordFile string) {
	os.WriteFile(recordFile, []byte(fmt.Sprintf("%v", effort)), 0644)
}
