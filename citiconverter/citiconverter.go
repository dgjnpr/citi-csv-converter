package citiconverter

import (
	"encoding/csv"
	"io"
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

// CitiIngest reads the Citi formated CSV file
func CitiIngest(r io.Reader) (*[][]string, error) {
	data := csv.NewReader(r)

	rows, err := data.ReadAll()
	if err != nil {
		return nil, err
	}

	return &rows, nil
}

// YnabHeaders is the header format for YNAB CSV file format
var YnabHeaders = []string{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"}

// ToYnab converts Citi CSV format to YNAB CSV forma
func ToYnab(citi *[][]string) *[][]string {
	(*citi)[0] = YnabHeaders

	// read all but the first line (which contain Citi headers)
	for i, r := range (*citi)[1:] {
		_inflow := ""
		_outflow := ""

		switch strings.Contains(r[BillingAmount], "-") {
		case true:
			_inflow = r[BillingAmount][2:]
		default:
			_outflow = r[BillingAmount]
		}
		(*citi)[i+1] = []string{r[TransactionDate], r[TransactionDetail], "Job Expense", "", _outflow, _inflow}

	}

	return citi
}
