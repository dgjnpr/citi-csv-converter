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
	data := csv.NewReader(r)
	var output [][]string

	data.Read()
	output = append(output, []string{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"})

	for {
		record, err := data.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, []string{record[2], record[5], "Job Expense", "", record[6], ""})
	}

	return output
}
