package dtdiff

import (
	"testing"
	"time"
)

var (
	dura, _ = time.ParseDuration("35h34m12s231ns")
	f       = NewFormatter(false, dura)
	fwu     = NewFormatter(true, dura)
)

func TestFormatLong(t *testing.T) {
	expected := f.Long()
	if expected != "1 days 11 hours 34 minutes 12 seconds\n" {
		t.Fatalf("invalid format %s", expected)
	}
}

func TestFormatShort(t *testing.T) {
	expected := f.Short()
	if expected != "1d11h34m12s\n" {
		t.Fatalf("invalid format %s", expected)
	}
}

func TestFormatHours(t *testing.T) {
	t.Run("with time unit", func(t *testing.T) {
		expected := fwu.Hours()
		if expected != "35.57 hours\n" {
			t.Fatalf("invalid format %s", expected)
		}
	})
	t.Run("without time unit", func(t *testing.T) {
		expected := f.Hours()
		if expected != "35.57\n" {
			t.Fatalf("invalid format %s", expected)
		}
	})
}

func TestFormatMinutes(t *testing.T) {
	t.Run("with time unit", func(t *testing.T) {
		expected := fwu.Minutes()
		if expected != "2134 minutes\n" {
			t.Fatalf("invalid format %s", expected)
		}
	})
	t.Run("without time unit", func(t *testing.T) {
		expected := f.Minutes()
		if expected != "2134\n" {
			t.Fatalf("invalid format %s", expected)
		}
	})
}

func TestFormatSecondst(t *testing.T) {
	t.Run("with time unit", func(t *testing.T) {
		expected := fwu.Seconds()
		if expected != "128052 seconds\n" {
			t.Fatalf("invalid format %s", expected)
		}
	})
	t.Run("without time unit", func(t *testing.T) {
		expected := f.Seconds()
		if expected != "128052\n" {
			t.Fatalf("invalid format %s", expected)
		}
	})
}

func TestFormatNanoseconds(t *testing.T) {
	t.Run("with time unit", func(t *testing.T) {
		expected := fwu.Nanoseconds()
		if expected != "128052000000231 nanoseconds\n" {
			t.Fatalf("invalid format %s", expected)
		}
	})
	t.Run("without time unit", func(t *testing.T) {
		expected := f.Nanoseconds()
		if expected != "128052000000231\n" {
			t.Fatalf("invalid format %s", expected)
		}
	})
}
