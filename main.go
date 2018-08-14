package main

import (
	"fmt"
	"os"
)

func main() {

	inFile := os.Args[1]
	if len(os.Args) > 2 {
		//		outFile := os.Args[2]
	} else {
		//		outFile := os.Args[1]
	}

	ledgerfile, err := os.Open(inFile)
	if err != nil {
		fmt.Printf("File %s not found", inFile)
		return
	}

	transactions, err := Parse(ledgerfile)
	if err != nil {
		fmt.Print(err)
	}
	sorted := Sort(transactions)
	fmt.Print(sorted)
	ledgerfile.Close()
}
