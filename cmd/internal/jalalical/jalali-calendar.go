package jalalical

import (
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

func FormatAsJalaliDay(t time.Time) string {
	pt := ptime.New(t)
	return pt.Format("yyyy MMM dd")
}

func TehranTimezone() *time.Location {
	return ptime.Iran()
}
