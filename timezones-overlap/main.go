package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func main() {

	// These four values are the sole configuration
	// Could have put them in a file or a JSON, but is it worth the hassle?
	ourTimezone := "Europe/Athens"
	theirTimezone := "US/Pacific"
	startTime := 9
	endTime := 18
    var stringHere, stringThere string

	localTimezone, err := time.LoadLocation(ourTimezone)
	foreignTimezone, err := time.LoadLocation(theirTimezone)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)

    // Headers for the table
	fmt.Fprintf(w, "\n\t%v \t\n", now.Format("2006-01-02"))
	fmt.Fprintf(w, "\t%v \t%v\t\n", localTimezone, foreignTimezone)

	for i := 0; i < 24; i++ {
		clockHere := time.Date(now.Year(), now.Month(), now.Day(), i, 0, 0, 0, localTimezone).In(localTimezone)    //.Format("15:04")
		clockThere := time.Date(now.Year(), now.Month(), now.Day(), i, 0, 0, 0, localTimezone).In(foreignTimezone) //.Format("15:04")
		stringHere = clockHere.Format("15:04")
		stringThere = clockThere.Format("15:04")

		if clockHere.Hour() >= startTime && clockHere.Hour() <= endTime {
			stringHere = "•" + stringHere
		}
		if clockThere.Hour() >= startTime && clockThere.Hour() <= endTime {
			stringThere = "•" + stringThere
		}
		fmt.Fprintf(w, "\t%v \t %v\t\n", stringHere, stringThere)
	}
	w.Flush()
	fmt.Println("")
}
