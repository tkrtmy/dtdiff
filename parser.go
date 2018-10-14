package dtdiff

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

// Parser parse date & time
type Parser interface {
	Parse(str string) (time.Time, error)
}

type parser struct {
	now time.Time
}

// ParseError is used when parse failed
type ParseError struct {
	datestr string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("Could not parse format for %q", e.datestr)
}

// NewParser returns a new parser
func NewParser(now time.Time) Parser {
	return &parser{now}
}

func (p *parser) Parse(str string) (time.Time, error) {
	t, err := dateparse.ParseLocal(str)
	if t.IsZero() {
		t, err = p.parseTime(str)
	}
	return t, err
}

func (p *parser) parseTime(str string) (time.Time, error) {
	// https://regexper.com/#%5E%280%3F%5B0-9%5D%7C1%5B0-9%5D%7C2%5B0-3%5D%29%3A%280%3F%5B0-9%5D%7C%5B1-5%5D%5B0-9%5D%29%24
	hhmm := regexp.MustCompile(`^(0?[0-9]|1[0-9]|2[0-3]):(0?[0-9]|[1-5][0-9])$`)
	// https://regexper.com/#%5E%280%3F%5B0-9%5D%7C1%5B0-9%5D%7C2%5B0-3%5D%29%3A%280%3F%5B0-9%5D%7C%5B1-5%5D%5B0-9%5D%29%3A%280%3F%5B0-9%5D%7C%5B1-5%5D%5B0-9%5D%29%24
	hhmmss := regexp.MustCompile(`^(0?[0-9]|1[0-9]|2[0-3]):(0?[0-9]|[1-5][0-9]):(0?[0-9]|[1-5][0-9])$`)
	switch {
	case hhmm.MatchString(str):
		return p.makeTodayDateTime(str), nil
	case hhmmss.MatchString(str):
		return p.makeTodayDateTime(str), nil
	}
	return time.Time{}, ParseError{str}
}

func (p *parser) makeTodayDateTime(str string) time.Time {
	year, month, day := p.now.Date()
	ss := strings.Split(str, ":")
	hour, _ := strconv.Atoi(ss[0])
	minute, _ := strconv.Atoi(ss[1])
	second := 0
	if len(ss) > 2 {
		second, _ = strconv.Atoi(ss[2])
	}
	return time.Date(year, month, day, hour, minute, second, 0, time.Local)
}
