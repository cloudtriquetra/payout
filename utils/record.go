package utils

import (
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

func (effort EffortEntry) RecordEffortToFile(recordFile string) {
	fmt.Println(effort.EmployeeName, effort.StartTimeStamp)
	os.WriteFile(recordFile, []byte(fmt.Sprintf("%v", effort)), 0644)
}
