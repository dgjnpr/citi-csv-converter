package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

const (
	// AccountNumber is the redacted CCard number
	AccountNumber = iota
	// AccountName is the CCard holder name
	AccountName
	// TransactionDate is the date the transaction took place
	TransactionDate
	// PostDate is the date the transaction cleared
	PostDate
	// ReferenceNumber is a transaction ID
	ReferenceNumber
	// TransactionDetail usually contains vendor name and location
	TransactionDetail
	// BillingAmount is the transaction amount in the CCard local currency
	BillingAmount
	// SourceCurrency is a three letter currency ID
	SourceCurrency
	// SourceAmount is the transaction amount in SourceCurrency
	SourceAmount
	// CustomerRef so far has been an empty field
	CustomerRef
	// EmployeeID is a 9 digit number. For US employees it's their SSN
	EmployeeID
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

	rows, err := data.ReadAll()
	if err != nil {
		return nil, err
	}

	// read all but the first line (which contain Citi headers)
	for _, r := range rows[1:] {
		switch strings.Contains(r[BillingAmount], "-") {
		case true:
			inflow := r[BillingAmount][2:]
			output = append(output, []string{r[TransactionDate], r[TransactionDetail], "Job Expense", "", "", inflow})
		default:
			output = append(output, []string{r[TransactionDate], r[TransactionDetail], "Job Expense", "", r[BillingAmount], ""})
		}
	}

	return output, nil
}
