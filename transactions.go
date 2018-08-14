package main

import (
	date "github.com/araddon/dateparse"
	"os"
	"time"
)

type Transaction struct {
	time time.Time
	text string
}

func GetTransaction(reader io.Reader) transaction {

	scanner := bufio.NewScanner(reader)
	transactionComplete := false
	var inTransaction bool
	var trans transaction

	for !transactinComplete {

		// Seclude the data form the line read in.
		line = scanner.Text()
		trimmedLine := strings.Trim(line, "\t")

		// Write comments unchanged
		if strings.HasPrefix(trimmedLine, ";") {
			trans.text += line + "\n"

			// Handle date parsing
		} else if !inTransaction {

			splitLine := strings.SplitN(line, " ", 2)
			rawDate := splitLine[0]

			// Record date
			time, err := dateparse.ParseAny(rawDate)
			if err != nil {
				return rtrn, fmt.Errorf("%d: Unable to parse date %s% %v\n", linecount, rawDate, err)
			}
			fmtDate := strings.Split(time.String(), " ")[0]

			trans.text += line + "\n"
			inTransaction = true

		} else if len(trimmedLine) == 0 {
			inTransaction = false
			rtrn += "\n"
			transactionComplete = true
		} else {
			rtrn += line + "\n"
		}
	}
	return *trans
}

// return the difference in time between the two trnasactions.
// Returns positive if the calling transaction takes place after the parameter transaction
func (this *Transaction) Compare(that Transaction) int64 {

	return this.Time.Unix() - that.Time.Unix()

}

func Parse(file *File) []Transaction {
	err = nil
	trans := []Transaction{}

	for err == nil {
		t, err := GetTransaction(file)
		if err != nil {
			trans = append(trans, t)
		}
	}
	return trans
}
