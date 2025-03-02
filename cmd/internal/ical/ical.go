package ical

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
)

type ical struct {
	cal *ics.Calendar
}

func NewIcal() *ical {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodPublish)
	return &ical{cal}
}

func (i *ical) AddEvent(day *time.Time, title string) {
	hash := md5.Sum([]byte(title))
	id := hex.EncodeToString(hash[:])
	event := i.cal.AddEvent(fmt.Sprintf("%s@%s",
		"github.com/raaminz/jalali-ical/1", id))
	date := day.UTC().Truncate(24 * time.Hour)
	event.SetCreatedTime(date)
	event.SetDtStampTime(date)
	event.SetAllDayStartAt(date)
	event.SetAllDayEndAt(date.Add(time.Hour * 24))
	event.SetSummary(title)
}

func (i *ical) Serialize() string {
	return i.cal.Serialize()
}
