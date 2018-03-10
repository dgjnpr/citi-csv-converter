package main

import (
	"encoding/csv"
	"log"
	"os"

	cc "github.com/dgjnpr/citi-csv-converter/citiconverter"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("cannot read csv: %v", err)
	}
	defer data.Close()

	citi, err := cc.CitiIngest(data)
	if err != nil {
		log.Fatalf("couldn't parse csv file: %v", err)
	}

	ynab := cc.ToYnab(citi)

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(*ynab)

	if err = w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}
