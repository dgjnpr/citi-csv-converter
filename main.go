package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("cannot read csv: %v", err)
	}
	defer data.Close()

	output := YnabParser(data)

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(output)

	if err = w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

// YnabParser blah blah blah
func YnabParser(r io.Reader) [][]string {
	var output [][]string
	output = append(output, []string{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"})

	data := csv.NewReader(r)
	// pop the headers from Citi
	data.Read()

	records, err := data.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		output = append(output, []string{record[2], record[5], "Job Expense", "", record[6], ""})
	}

	return output
}
