package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	inFile, err := os.Open("/Users/dgethings/Downloads/Statement.csv")
	if err != nil {
		log.Fatalf("cannot read csv: %v", err)
	}
	defer inFile.Close()

	data := csv.NewReader(inFile)

	w := csv.NewWriter(os.Stdout)

	headers, err := data.Read()
	if err != nil {
		log.Fatal(err)
	}
	w.Write(headers[0:])
	w.Write([]string{headers[2], headers[6]})

	for {
		record, err := data.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]string{record[2], record[6]})
	}

	w.Flush()
}
