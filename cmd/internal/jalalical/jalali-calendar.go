package jalalical

import (
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

type jalaliCal struct {
	pt ptime.Time
}

func NewJalaliCal(t time.Time) *jalaliCal {
	return &jalaliCal{ptime.New(t)}
}

func (j *jalaliCal) FormatDay() string {
	return j.pt.Format("yyyy/MM/dd (MMM)")
}

func TehranLocation() *time.Location {
	return ptime.Iran()

}

// Year is not used method
func (j *jalaliCal) Year() int {
	return j.pt.Year()

}
