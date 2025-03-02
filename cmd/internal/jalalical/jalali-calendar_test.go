package jalalical_test

import (
	"testing"
	"time"

	"ramin.tech/cmd/jalai-ical/cmd/internal/jalalical"
)

func TestMain(t *testing.T) {
	var r time.Time = time.Date(2016, time.January, 1, 12, 1, 1, 0,
		jalalical.TehranTimezone())
	got := jalalical.FormatAsJalaliDay(r)
	const want = "1394 دی 11"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
