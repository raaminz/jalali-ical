package ical_test

import (
	"strings"
	"testing"
	"time"

	"ramin.tech/cmd/jalai-ical/cmd/internal/ical"
)

const expected = `
BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//arran4//Golang ICS Library
METHOD:PUBLISH
BEGIN:VEVENT
UID:github.com/raaminz/jalali-ical/1@5d41402abc4b2a76b9719d911017c592
CREATED:20221225T000000Z
DTSTAMP:20221225T000000Z
DTSTART;VALUE=DATE:20221225
DTEND;VALUE=DATE:20221226
SUMMARY:hello
END:VEVENT
BEGIN:VEVENT
UID:github.com/raaminz/jalali-ical/1@4ccefd5121d7d991c5d663658225dd39
CREATED:20221226T000000Z
DTSTAMP:20221226T000000Z
DTSTART;VALUE=DATE:20221226
DTEND;VALUE=DATE:20221227
SUMMARY:1394/10/11 (دی)
END:VEVENT
END:VCALENDAR
`

func TestCreate(t *testing.T) {
	today := time.Date(2022, 12, 25, 0, 0, 0, 0, time.UTC)
	ical := ical.NewIcal()
	ical.AddEvent(&today, "hello")
	tomorrow := today.Add(24 * time.Hour)
	ical.AddEvent(&tomorrow, "1394/10/11 (دی)")
	got := ical.Serialize()
	if strings.TrimSpace(got) != strings.TrimSpace(expected) {
		t.Errorf("got %s want %s", got, expected)
	}
}
