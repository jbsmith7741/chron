package chron

import (
	"fmt"
	"strconv"
	"time"
)

const (
	DashDDMMMYY       = "02-Jan-06"
	DashDDMMMYYYY     = "02-Jan-2006"
	DashMMDDYY        = "01-02-06"
	DashMMDDYYYY      = "01-02-2006"
	SlashMMDDYY       = "01/02/06"
	SlashMMDDYYY      = "01/02/2006"
	ShortDateTime     = SlashMMDDYYY + " 15:04 PM"
	ShortSecond       = SlashMMDDYYY + " 15:04:05 PM"
	ShortDateTime24   = SlashMMDDYYY + " 15:04"
	ShortSecond24     = SlashMMDDYYY + " 15:04:05"
	DashMonth         = "Jan-2006"
	DashMonthShort    = "Jan-06"
	DashNumMonth      = "01-2006"
	DashNumMonthShort = "01-06"
	CCMonth           = "01/06"
	YearFmt           = "2006"
)

var ParseFormats = []string{
	time.RFC3339,
	time.RFC3339Nano,
	time.RFC822,
	time.RFC822Z,
	time.UnixDate,
	time.ANSIC,
	time.RubyDate,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC850,
	DashDDMMMYY,
	DashDDMMMYYYY,
	DashMMDDYY,
	DashMMDDYYYY,
	SlashMMDDYY,
	SlashMMDDYYY,
	ShortDateTime,
	ShortSecond,
	ShortDateTime24,
	ShortSecond24,
	DashMonth,
	DashMonthShort,
	DashNumMonth,
	DashNumMonthShort,
	CCMonth,
	YearFmt,
}

var ParseFunctions = []func(string) (time.Time, error){
	ParseWithFormats,
	ParseUnixSeconds,
}

func ParseUnixSeconds(secs string) (time.Time, error) {
	i, err := strconv.Atoi(secs)
	if err != nil {
		return ZeroTime(), err
	}
	return time.Unix(int64(i), 0).UTC(), nil
}

func ParseWithFormats(s string) (time.Time, error) {
	for _, layout := range ParseFormats {
		t, err := time.Parse(layout, s)
		if err != nil {
			continue
		}
		return t, nil
	}
	return ZeroTime(), fmt.Errorf("string didn't match an attempted format")
}
