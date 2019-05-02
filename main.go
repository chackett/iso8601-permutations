package main

import (
	"fmt"
	"os"
	"time"
)
//https://en.wikipedia.org/wiki/ISO_8601
//https://www.iso.org/iso-8601-date-and-time-format.html
//2006-01-02T15:04:05Z07:00

var dates = []string{
	"2006-01-02",
	"2006-01",
	//"--01-02",
	"20060102",
	//"--0201",
}
var times = []string{
	"150405.999",
	"150405",
	"1504",
	"15",
	"15:04:05.99.999",
	"15:04:05",
	"15:04",
}
var zones = []string{
	"Z",
	"+07:00",
	"+0700",
	"+07",
	"-07:00",
	"-0700",
	"-07",
}

var validFormats = constructValidFormats()

func constructValidFormats() []string {
	var result []string
	for _, date := range dates {
		for _, time := range times {
			for _, zone := range zones {
				result = append(result, fmt.Sprintf("%sT%s%s", date, time, zone))
			}
		}
	}

	return result
}

func main() {
	perms, err := permutateDates(time.Now(), validFormats)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, v := range perms {
		fmt.Println(v)
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Printf("%d valid formats\n", len(perms))
}

func permutateDates(dt time.Time, validFormats []string) ([]string, error) {
	var result []string
	for _, v := range validFormats {
		result = append(result, dt.Format(v))
	}

	return result, nil
}

