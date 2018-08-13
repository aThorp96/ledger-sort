package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	dateparse "github.com/araddon/dateparse"
)

/*
Get Demarkation returns which of the following is used to
seperate each part of a date.
`-`, `.`, or `/`
*/
func GetDemarkation(date string) string {
	index := -1
	// Tests for slashes
	index = strings.Index(date, "/")
	if index >= 0 {
		return date[index : index+1]
	}
	//Tests for dashes
	index = strings.Index(date, "-")
	if index >= 0 {
		return date[index : index+1]
	}

	//Tests for periods
	index = strings.Index(date, ".")
	if index >= 0 {
		return date[index : index+1]
	}
	return ""
}

func Parse(reader io.Reader) (string, error) {

	scanner := bufio.NewScanner(reader)
	rtrn := ""
	curYear := ""
	linecount := 0
	var line string
	var inTransaction bool

	for scanner.Scan() {

		// Seclude the data form the line read in.
		line = scanner.Text()
		trimmedLine := strings.Trim(line, "\t")
		log.Printf("Reading line %d\n", linecount)
		linecount++

		// Write comments unchanged
		if strings.HasPrefix(trimmedLine, ";") {
			rtrn += line + "\n"
			log.Printf("Line %d is a comment\n", linecount-1)
			// Handle date parsing
		} else if !inTransaction {

			// Splits the line apart by the first space.
			// Seperates the date for correction
			splitLine := strings.SplitN(line, " ", 2)
			rawDate := splitLine[0]
			log.Printf("Extracted raw date %s from line %d", rawDate, linecount-1)

			// Append the current year if not present.
			if len(rawDate) < 8 {
				if len(curYear) == 0 {
					return rtrn, fmt.Errorf("%d: Unable to correct date. No fallback year\n", linecount)
				} else {
					demarkation := GetDemarkation(rawDate)
					rawDate = curYear + demarkation + rawDate
				}
			} else {
				demarkation := GetDemarkation(rawDate)
				curYear = strings.Split(rawDate, demarkation)[0]
			}
			// Convert Date to yyyy-mm-dd
			time, err := dateparse.ParseAny(rawDate)
			if err != nil {
				return rtrn, fmt.Errorf("%d: Unable to parse date %s% %v\n", linecount, rawDate, err)
			}
			fmtDate := strings.Split(time.String(), " ")[0]

			rtrn += fmtDate + " " + splitLine[1] + "\n"
			inTransaction = true

		} else if len(trimmedLine) == 0 {
			inTransaction = false
			rtrn += "\n"
		} else {
			rtrn += line + "\n"
		}
	}
	return rtrn, nil
}
