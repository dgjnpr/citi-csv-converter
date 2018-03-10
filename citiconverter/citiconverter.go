package citiconverter

import (
	"encoding/csv"
	"io"
	"strings"
)

const (
	// AccountNumber is the redacted CCard number
	accountNumber = iota
	// AccountName is the CCard holder name
	accountName
	// TransactionDate is the date the transaction took place
	transactionDate
	// PostDate is the date the transaction cleared
	postDate
	// ReferenceNumber is a transaction ID
	referenceNumber
	// TransactionDetail usually contains vendor name and location
	transactionDetail
	// BillingAmount is the transaction amount in the CCard local currency
	billingAmount
	// SourceCurrency is a three letter currency ID
	sourceCurrency
	// SourceAmount is the transaction amount in SourceCurrency
	sourceAmount
	// CustomerRef so far has been an empty field
	customerRef
	// EmployeeID is a 9 digit number. For US employees it's their SSN
	employeeID
)

// CitiIngest reads the Citi formated CSV file
// probably should validate this data, but the only thing we care about is the
// column order. I doubt that will change much
func CitiIngest(r io.Reader) (*[][]string, error) {
	data := csv.NewReader(r)

	rows, err := data.ReadAll()
	if err != nil {
		return nil, err
	}

	return &rows, nil
}

var ynabHeaders = []string{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"}

// ToYnab converts Citi CSV format to YNAB CSV format
func ToYnab(citi *[][]string) *[][]string {
	(*citi)[0] = ynabHeaders

	// read all but the first line (which contain headers)
	for i, r := range (*citi)[1:] {
		_inflow := ""
		_outflow := ""

		switch strings.Contains(r[billingAmount], "-") {
		case true:
			_inflow = r[billingAmount][2:]
		default:
			_outflow = r[billingAmount]
		}

		// rewrite current row to YNAB format
		// index off by -1 as range starts at 1
		(*citi)[i+1] = []string{
			r[transactionDate],
			r[transactionDetail],
			"Job Expense",
			"",
			_outflow,
			_inflow,
		}

	}

	return citi
}
