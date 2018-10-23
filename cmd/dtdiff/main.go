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

	f := dtdiff.NewFormatter(!quiet, dura)
	var msg string
	switch {
	case sht:
		msg = f.Short()
	case hour:
		msg = f.Hours()
	case min:
		msg = f.Minutes()
	case sec:
		msg = f.Seconds()
	case nanosec:
		msg = f.Nanoseconds()
	default:
		msg = f.Long()
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
