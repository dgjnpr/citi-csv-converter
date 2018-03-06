package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("cannot read csv: %v", err)
	}
	defer data.Close()

	output, err := YnabParser(data)
	if err != nil {
		log.Fatalf("couldn't parse csv file: %v", err)
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(output)

	if err = w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

// YnabParser blah blah blah
func YnabParser(r io.Reader) ([][]string, error) {
	var output [][]string
	output = append(output, []string{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"})

	data := csv.NewReader(r)

	records, err := data.ReadAll()
	if err != nil {
		return nil, err
	}

	// read all but the first line (which contain Citi headers)
	for _, r := range records[1:] {
		switch strings.Contains(r[6], "-") {
		case true:
			inflow := r[6][2:]
			output = append(output, []string{r[2], r[5], "Job Expense", "", "", inflow})
		default:
			output = append(output, []string{r[2], r[5], "Job Expense", "", r[6], ""})
		}
	}

	return output, nil
}
