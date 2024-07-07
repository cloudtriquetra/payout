package utils

import (
	"strings"
	"time"
)

func GetDuration(startDateTime time.Time, endDateTime time.Time) time.Duration {

	duration := endDateTime.Sub(startDateTime)
	return duration

}

func GetTimeStamp(date string, eventTime string) time.Time {
	const dateTimeFormat = "2-01-2006 15:04"
	tz, _ := time.LoadLocation("Europe/Warsaw")
	var timeStampString strings.Builder

	timeStampString.WriteString(date)
	timeStampString.WriteString(" ")
	timeStampString.WriteString(eventTime)
	timeStamp, _ := time.Parse(dateTimeFormat, timeStampString.String())
	return timeStamp.In(tz)

}
