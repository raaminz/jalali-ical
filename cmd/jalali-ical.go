package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"ramin.tech/cmd/jalai-ical/cmd/internal/ical"
	"ramin.tech/cmd/jalai-ical/cmd/internal/jalalical"
)

var (
	year = flag.Int("year", time.Now().Year(), "Gregorian Year to be converted")
)

func main() {
	flag.Parse()

	if err := Main(*year); err != nil {
		log.Fatal(err)
	}
}

func Main(year int) error {
	fmt.Printf("generate the persian ical file for year %d", year)

	firstDay := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	ical := ical.NewIcal()
	currday := &firstDay
	for year == currday.Year() {
		day := *currday
		jalali := jalalical.NewJalaliCal(day)
		ical.AddEvent(day, jalali.FormatDay())
		day = currday.Add(24 * time.Hour)
		year = currday.Year()
		currday = &day
		log.Println(jalali.FormatDay())
		log.Println(year)
	}
	iclContent := ical.Serialize()
	if iclContent == "" {
		panic("something went wrong")
	}

	firstJalaliYear := jalalical.NewJalaliCal(firstDay)
	secondJalaliYear := jalalical.NewJalaliCal(*currday)
	// persian (jalali) calendar, filename
	fileName := fmt.Sprintf("%d_(%d-%d)_persian_calendar.ics",
		year, firstJalaliYear.Year(), secondJalaliYear.Year())

	return writeToFile(fileName, iclContent)
}

func writeToFile(fileName string, data string) error {
	log.Println("writing to file " + fileName)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
