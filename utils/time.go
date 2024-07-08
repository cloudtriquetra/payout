package utils

import (
	"errors"
	"strings"
	"time"
)

func GetDuration(startDateTime time.Time, endDateTime time.Time) time.Duration {

	duration := endDateTime.Sub(startDateTime)
	return duration

}

func GetTimeStamp(date string, eventTime string) (time.Time, error) {
	const dateTimeFormat = "2-01-2006 15:04"
	tz, _ := time.LoadLocation("Europe/Warsaw")
	var timeStampString strings.Builder

	timeStampString.WriteString(date)
	timeStampString.WriteString(" ")
	timeStampString.WriteString(eventTime)
	timeStamp, err := time.ParseInLocation(dateTimeFormat, timeStampString.String(), tz)
	if err != nil {
		return timeStamp, errors.New("invalid date or time format entered")
	}
	return timeStamp, nil

}
