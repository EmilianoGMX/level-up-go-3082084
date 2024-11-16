package main

import (
	"flag"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	parsed, err := time.Parse(expectedFormat, target)
	if err != nil {
		panic(err)
	}

	return parsed
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	duration := time.Now().Sub(target)

	return duration.Hours() / 24
}

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
