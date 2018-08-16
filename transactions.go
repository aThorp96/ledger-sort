package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	date "github.com/araddon/dateparse"
)

type Transaction struct {
	time time.Time
	text string
}

func GetTransaction(reader io.Reader) (Transaction, error) {

	scanner := bufio.NewScanner(reader)
	transactionComplete := false
	var inTransaction bool
	var trans Transaction

	for !transactionComplete && err := scanner.Err() != nil {

		// Seclude the data form the line read in.
		scanner.Scan()
		line := scanner.Text()
		fmt.Println("line: ", line)
		trimmedLine := strings.Trim(line, "\t")

		// Write comments unchanged
		if strings.HasPrefix(trimmedLine, ";") {
			trans.text += line + "\n"

		} else if len(trimmedLine) == 0 {
			inTransaction = false
			trans.text += "\n"
			transactionComplete = true
			// Handle date parsing
		} else if !inTransaction {

			splitLine := strings.SplitN(line, " ", 2)
			rawDate := splitLine[0]

			// Record date
			time, err := date.ParseAny(rawDate)
			if err != nil {
				return trans, fmt.Errorf("Unable to parse date %s% %v\n", rawDate, err)
			}
			trans.time = time

			trans.text += line + "\n"
			inTransaction = true

		} else {
			trans.text += line + "\n"
		}
		fmt.Println(line) //log
	}
	return trans, nil
}

// return the difference in time between the two trnasactions.
// Returns positive if the calling transaction takes place after the parameter transaction
func (this *Transaction) Compare(that Transaction) int64 {

	return this.time.Unix() - that.time.Unix()

}

func Parse(file *os.File) ([]Transaction, error) {
	fmt.Println("Parsing") // log
	var err error
	var t Transaction
	trans := []Transaction{}

	for err == nil {
		fmt.Println("running loop") //log
		t, err = GetTransaction(file)
		if err != nil {
			trans = append(trans, t)
		}
	}
	return trans, err
}
