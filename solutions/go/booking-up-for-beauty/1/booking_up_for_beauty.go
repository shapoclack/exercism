package booking

import (
	"fmt"
	"time"
)

// HasPassed returns whether a date has passed.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

func HasPassed(date string) bool {
	now := time.Now()
	layout := "January 2, 2006 15:04:05"
	act, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	if now.After(act) {
		return true
	} else if now.Before(act) {
		return false
	} else {
		return false
	}
}

func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    act, err := time.Parse(layout, date)
    if err != nil {
        return false
    }
    return act.Hour() >= 12 && act.Hour() < 18
}
func Description(date string) string {
	meet := Schedule(date)
	description := fmt.Sprintf("You have an appointment on %s, at %s.", meet.Format("Monday, January 2, 2006"), meet.Format("15:04"))
	return description
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}