package dtdiff

import (
	"reflect"
	"testing"
	"time"
)

type dataSet struct {
	in, out string
}

// Parse logic to be secured araddon/dateparse.
// So this test cover focus on hh:mm:ss, hh:mm formats
var testInput = []dataSet{
	{in: "21:34:32", out: "2018-10-14 21:34:32"},
	{in: "0:0:0", out: "2018-10-14 00:00:00"},
	{in: "1:2:3", out: "2018-10-14 01:02:03"},
	{in: "00:00:00", out: "2018-10-14 00:00:00"},
	{in: "23:59:59", out: "2018-10-14 23:59:59"},
}

func TestParseSucceed(t *testing.T) {
	now := time.Date(2018, 10, 14, 12, 12, 12, 12, time.Local)
	p := NewParser(now)
	for _, ti := range testInput {
		time, err := p.Parse(ti.in)
		if err != nil {
			t.Fatalf("test failed %s\n", err.Error())
		}
		if time.Format("2006-01-02 15:04:05") != ti.out {
			t.Fatalf("expected: %s, actual: %s\n", ti.out, time.Format("2006-01-02 15:04:05"))
		}
	}
}

var testError = []dataSet{
	{in: "21::32"},
	{in: "12:61"},
	{in: "12:00:61"},
	{in: "0:0:123"},
	{in: "-00:00:00"},
	{in: "24:00:00"},
	{in: "AA:BB:CC"},
}

func TestParseError(t *testing.T) {
	now := time.Date(2018, 10, 14, 12, 12, 12, 12, time.Local)
	p := NewParser(now)
	for _, te := range testError {
		_, err := p.Parse(te.in)
		if err == nil {
			t.Fatalf("test failed %s\n", te.in)
		}
		switch err.(type) {
		case ParseError:
			// ok
		default:
			t.Fatalf("error type is %s\n", reflect.TypeOf(err).String())
		}
	}
}
