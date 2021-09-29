package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jilleJr/go-timetrap/pkg/timetrap"
)

func main() {
	config, err := timetrap.NewConfigLocal()
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
	if config.DatabaseFile == "" {
		fmt.Println("err: no database file specified in config:", timetrap.DefaultConfigPath)
		os.Exit(1)
	}

	db, err := timetrap.NewDB(config.DatabaseFile)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	currentSheet, err := db.GetCurrentSheet()
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	lastCheckoutID, err := db.GetLastCheckoutID()
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	activeEntry, err := db.GetActiveEntry()
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	entriesToday, err := db.GetEntriesTimeRange(today, today.Add(24*time.Hour))
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	fmt.Println("current sheet:   ", currentSheet)
	fmt.Println("last checkout ID:", lastCheckoutID)
	fmt.Println("active entry:    ", activeEntry)

	fmt.Println("entries today:")
	var sum time.Duration
	for _, entry := range entriesToday {
		fmt.Println("  -", entry)
		sum += entry.Duration()
	}
	if len(entriesToday) == 0 {
		fmt.Println("  *none*")
	}

	fmt.Println("time today:      ", sum)

	dayLength := time.Hour * time.Duration(config.DayLengthHours)
	fmt.Println("time today %:    ", int64(100*sum/dayLength), "%")
}
