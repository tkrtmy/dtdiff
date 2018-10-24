package dtdiff

import (
	"fmt"
	"time"
)

// Formatter is time diff put into words
type Formatter struct {
	// without time unit?
	withoutUnit bool
	// time diff duration
	duration time.Duration
}

// NewFormatter returns new Formatter
func NewFormatter(withoutUnit bool, duration time.Duration) *Formatter {
	return &Formatter{withoutUnit, duration}
}

func resolve(duration time.Duration) (days, hours, mins, secs int) {
	whole := int(duration.Hours())
	days = whole / 24
	hours = whole % 24
	mins = int(duration.Minutes()) % 60
	secs = int(duration.Seconds()) % 60
	return
}

// Long description (days, hours, minutes, seconds) of time diff
func (f *Formatter) Long() string {
	days, hours, mins, secs := resolve(f.duration)
	return fmt.Sprintf("%d days %d hours %d minutes %d seconds\n", days, hours, mins, secs)
}

// Short description (days, hours, minutes, seconds) of time diff
func (f *Formatter) Short() string {
	days, hours, mins, secs := resolve(f.duration)
	return fmt.Sprintf("%dd%dh%dm%ds\n", days, hours, mins, secs)
}

// Hours display hour, the second decimal. ex. 4.83 hours
func (f *Formatter) Hours() string {
	var format string
	if f.withoutUnit {
		format = "%.2f\n"
	} else {
		format = "%.2f hours\n"
	}
	return fmt.Sprintf(format, f.duration.Hours())
}

// Minutes display minutes. ex. 123 minutes
func (f *Formatter) Minutes() string {
	var format string
	if f.withoutUnit {
		format = "%d\n"
	} else {
		format = "%d minutes\n"
	}
	return fmt.Sprintf(format, int(f.duration.Minutes()))
}

// Seconds display seconds. ex. 65555 seconds
func (f *Formatter) Seconds() string {
	var format string
	if f.withoutUnit {
		format = "%d\n"
	} else {
		format = "%d seconds\n"
	}
	return fmt.Sprintf(format, int(f.duration.Seconds()))
}

// Nanoseconds display. ex 810821916 nanoseconds
func (f *Formatter) Nanoseconds() string {
	var format string
	if f.withoutUnit {
		format = "%d\n"
	} else {
		format = "%d nanoseconds\n"
	}
	return fmt.Sprintf(format, f.duration.Nanoseconds())
}
