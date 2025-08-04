package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// Format: "7/13/2020 20:32:00"
	const layout = "1/2/2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Printf("Error parsing date '%s': %v\n", date, err)
		return time.Time{}
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// Format: "October 3, 2019 20:32:00"
	const layout = "January 2, 2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Printf("Error parsing date '%s': %v\n", date, err)
		return false
	}
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// Format: "Friday, March 8, 1974 12:02:02"
	const layout = "Monday, January 2, 2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Printf("Error parsing date '%s': %v\n", date, err)
		return false
	}

	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	const parseLayout = "1/2/2006 15:04:05"

	t, err := time.Parse(parseLayout, date)
	if err != nil {
		return fmt.Sprintf("Error parsing date: %v", err)
	}

	const formatLayout = "Monday, January 2, 2006, at 15:04."

	// Format the parsed time `t` into the desired output string.
	return fmt.Sprintf("You have an appointment on %s", t.Format(formatLayout))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	// Get the current year.
	currentYear := time.Now().Year()

	anniversary := time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)

	return anniversary
}
