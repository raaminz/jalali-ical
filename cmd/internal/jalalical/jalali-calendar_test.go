package jalalical_test

import (
	"testing"
	"time"

	"ramin.tech/cmd/jalai-ical/cmd/internal/jalalical"
)

func TestMain(t *testing.T) {
	var r time.Time = time.Date(2016, time.January, 1, 12, 1, 1, 0,
		jalalical.TehranLocation())
	j := jalalical.NewJalaliCal(r)
	t.Run("format day", func(t *testing.T) {
		got := j.FormatDay()
		const want = "1394/10/11 (دی)"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("year", func(t *testing.T) {
		got := j.Year()
		const want = 1394
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
