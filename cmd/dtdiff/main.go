package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/tkrtmy/dtdiff"
)

var (
	sht     bool
	hour    bool
	min     bool
	sec     bool
	nanosec bool
	quiet   bool
	until   bool
	version bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: dtdiff [options] [datetime(from)] [datetime(to)]\n\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&sht, "short", false, "short description")
	flag.BoolVar(&hour, "h", false, "display hours only")
	flag.BoolVar(&min, "m", false, "display minutes only")
	flag.BoolVar(&sec, "s", false, "display seconds only")
	flag.BoolVar(&nanosec, "n", false, "display nanoseconds only")
	flag.BoolVar(&quiet, "q", false, "display without time unit")
	flag.BoolVar(&until, "until", false, "calculate until a given one")
	flag.BoolVar(&version, "v", false, "display current version")
	flag.Parse()
}

func run() error {
	if version {
		fmt.Fprintf(os.Stdout, "%s\n", Version)
		return nil
	}
	if len(flag.Args()) < 1 {
		return fmt.Errorf("You must give arguments at least one")
	}
	parser := dtdiff.NewParser(time.Now())
	ts := []time.Time{}
	for i, arg := range flag.Args() {
		t, err := parser.Parse(arg)
		if err != nil {
			return fmt.Errorf("The %d argument: %s", i+1, err.Error())
		}
		ts = append(ts, t)
	}
	var dura time.Duration
	if len(ts) == 1 {
		// give one augument, calculate diff between ts[0] and now.
		if until {
			// form now to ts[0]
			dura = dtdiff.CalculateDiffUntil(ts[0])
		} else {
			// form ts[0] to now
			dura = dtdiff.CalculateDiffSince(ts[0])
		}
	} else {
		dura = dtdiff.CalculateDiff(ts[0], ts[1])
	}

	var msg string
	switch {
	case sht:
		msg = short(dura)
	case hour:
		msg = hours(dura)
	case min:
		msg = minutes(dura)
	case sec:
		msg = seconds(dura)
	case nanosec:
		msg = nanoseconds(dura)
	default:
		msg = long(dura)
	}

	fmt.Fprintf(os.Stdout, msg)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func resolve(duration time.Duration) (days, hours, mins, secs int) {
	whole := int(duration.Hours())
	days = whole / 24
	hours = whole % 24
	mins = int(duration.Minutes()) % 60
	secs = int(duration.Seconds()) % 60
	return
}

func long(duration time.Duration) string {
	days, hours, mins, secs := resolve(duration)
	return fmt.Sprintf("%d days %d hours %d minutes %d seconds\n", days, hours, mins, secs)
}

func short(duration time.Duration) string {
	days, hours, mins, secs := resolve(duration)
	return fmt.Sprintf("%dd%dh%dm%ds\n", days, hours, mins, secs)
}

func hours(duration time.Duration) string {
	var format string
	if quiet {
		format = "%.2f\n"
	} else {
		format = "%.2f hours\n"
	}
	return fmt.Sprintf(format, duration.Hours())
}

func minutes(duration time.Duration) string {
	var format string
	if quiet {
		format = "%d\n"
	} else {
		format = "%d minutes\n"
	}
	return fmt.Sprintf(format, int(duration.Minutes()))
}

func seconds(duration time.Duration) string {
	var format string
	if quiet {
		format = "%d\n"
	} else {
		format = "%d seconds\n"
	}
	return fmt.Sprintf(format, int(duration.Seconds()))
}

func nanoseconds(duration time.Duration) string {
	var format string
	if quiet {
		format = "%d\n"
	} else {
		format = "%d nanoseconds\n"
	}
	return fmt.Sprintf(format, duration.Nanoseconds())
}
